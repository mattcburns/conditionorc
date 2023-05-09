package app

import (
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/metal-toolbox/conditionorc/internal/model"
	ptypes "github.com/metal-toolbox/conditionorc/pkg/types"
	"github.com/pkg/errors"
	"go.hollow.sh/toolbox/events"
)

var (
	ErrConfig = errors.New("configuration error")
)

// Configuration holds application configuration read from a YAML or set by env variables.
//
// nolint:govet // prefer readability over field alignment optimization for this case.
type Configuration struct {
	// file is the configuration file path
	file string `mapstructure:"-"`

	// LogLevel is the app verbose logging level.
	// one of - info, debug, trace
	LogLevel string `mapstructure:"log_level"`

	// ListenAddress is the server listen address
	ListenAddress string `mapstructure:"listen_address"`

	// Concurrency declares how many events the orchestrator will process concurrently.
	Concurrency int `mapstructure:"concurrency"`

	// StoreKind indicates the kind of store to store server conditions
	// supported parameter value - serverservice
	StoreKind model.StoreKind `mapstructure:"store_kind"`

	// ConditionDefinitions holds one or more condition definitions the conditionorc API, orchestrator support.
	ConditionDefinitions ptypes.ConditionDefinitions `mapstructure:"conditions"`

	// ServerserviceOptions defines the serverservice client configuration parameters
	//
	// This parameter is required when StoreKind is set to serverservice.
	ServerserviceOptions ServerserviceOptions `mapstructure:"serverservice"`

	// EventsBrokerKind indicates the kind of event broker configuration to enable,
	//
	// Supported parameter value - nats
	EventsBrokerKind model.EventBrokerKind `mapstructure:"events_broker_kind"`

	// NatsOptions defines the NATs events broker configuration parameters.
	//
	// This parameter is required when EventsBrokerKind is set to nats.
	NatsOptions events.NatsOptions `mapstructure:"nats"`
}

// ServerserviceOptions defines configuration for the Serverservice client.
// https://github.com/metal-toolbox/hollow-serverservice
type ServerserviceOptions struct {
	EndpointURL          *url.URL
	FacilityCode         string   `mapstructure:"facility_code"`
	Endpoint             string   `mapstructure:"endpoint"`
	OidcIssuerEndpoint   string   `mapstructure:"oidc_issuer_endpoint"`
	OidcAudienceEndpoint string   `mapstructure:"oidc_audience_endpoint"`
	OidcClientSecret     string   `mapstructure:"oidc_client_secret"`
	OidcClientID         string   `mapstructure:"oidc_client_id"`
	OidcClientScopes     []string `mapstructure:"oidc_client_scopes"`
	DisableOAuth         bool     `mapstructure:"disable_oauth"`
}

func (a *App) LoadConfiguration() error {
	a.v.SetConfigType("yaml")
	a.v.SetEnvPrefix(model.AppName)
	a.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	a.v.AutomaticEnv()

	fh, err := os.Open(a.Config.file)
	if err != nil {
		return errors.Wrap(err, ErrConfig.Error())
	}

	if err = a.v.ReadConfig(fh); err != nil {
		return errors.Wrap(err, ErrConfig.Error()+" "+a.Config.file)
	}

	if err := a.v.Unmarshal(a.Config); err != nil {
		return errors.Wrap(err, a.Config.file)
	}

	if err := a.envVarOverrides(); err != nil {
		return errors.Wrap(err, ErrConfig.Error())
	}

	return nil
}

func (a *App) envVarOverrides() error {
	if a.v.GetString("log.level") != "" {
		a.Config.LogLevel = a.v.GetString("log.level")
	}

	if a.v.GetString("app.kind") != "" {
		a.AppKind = model.AppKind(a.v.GetString("app.kind"))
	}

	if a.v.GetInt("concurrency") != 0 {
		a.Config.Concurrency = a.v.GetInt("concurrency")
	}

	if a.v.GetString("listen.address") != "" {
		a.Config.ListenAddress = a.v.GetString("listen.address")
	}

	if a.v.GetString("store.kind") != "" {
		a.Config.ListenAddress = a.v.GetString("store.kind")
	}

	if len(a.Config.ConditionDefinitions) == 0 {
		return errors.Wrap(ErrConfig, "expected one or more condition definitions")
	}

	if a.Config.EventsBrokerKind == model.NatsEventBroker {
		if err := a.envVarNatsOverrides(); err != nil {
			return err
		}
	}

	switch a.Config.StoreKind {
	case model.ServerserviceStore:
		if err := a.envVarServerserviceOverrides(); err != nil {
			return err
		}
	default:
		return errors.Wrap(ErrConfig, "no/unknown store kind parameter")
	}

	return nil
}

// NATs streaming configuration
var (
	defaultNatsConnectTimeout = 100 * time.Millisecond
)

func (a *App) envVarNatsOverrides() error {
	if a.v.GetString("nats.url") != "" {
		a.Config.NatsOptions.URL = a.v.GetString("nats.url")
	}

	if a.Config.NatsOptions.URL == "" {
		return errors.New("missing parameter: nats.url")
	}

	if a.v.GetString("nats.stream.user") != "" {
		a.Config.NatsOptions.StreamUser = a.v.GetString("nats.stream.user")
	}

	if a.v.GetString("nats.stream.pass") != "" {
		a.Config.NatsOptions.StreamPass = a.v.GetString("nats.stream.pass")
	}

	if a.v.GetString("nats.creds.file") != "" {
		a.Config.NatsOptions.CredsFile = a.v.GetString("nats.creds.file")
	}

	if a.v.GetString("nats.stream.name") != "" {
		a.Config.NatsOptions.Stream.Name = a.v.GetString("nats.stream.name")
	}

	if a.Config.NatsOptions.ConnectTimeout == 0 {
		a.Config.NatsOptions.ConnectTimeout = defaultNatsConnectTimeout
	}

	return nil
}

// Server service configuration options

// nolint:gocyclo // parameter validation is cyclomatic
func (a *App) envVarServerserviceOverrides() error {
	if a.v.GetString("serverservice.endpoint") != "" {
		a.Config.ServerserviceOptions.Endpoint = a.v.GetString("serverservice.endpoint")
	}

	if a.v.GetString("serverservice.facilitycode") != "" {
		a.Config.ServerserviceOptions.FacilityCode = a.v.GetString("serverservice.facilitycode")
	}

	if a.Config.ServerserviceOptions.FacilityCode == "" {
		return errors.New("serverservice facility code not defined")
	}

	endpointURL, err := url.Parse(a.Config.ServerserviceOptions.Endpoint)
	if err != nil {
		return errors.New("serverservice endpoint URL error: " + err.Error())
	}

	a.Config.ServerserviceOptions.EndpointURL = endpointURL

	if a.v.GetString("serverservice.disable.oauth") != "" {
		a.Config.ServerserviceOptions.DisableOAuth = a.v.GetBool("serverservice.disable.oauth")
	}

	if a.Config.ServerserviceOptions.DisableOAuth {
		return nil
	}

	if a.v.GetString("serverservice.oidc.issuer.endpoint") != "" {
		a.Config.ServerserviceOptions.OidcIssuerEndpoint = a.v.GetString("serverservice.oidc.issuer.endpoint")
	}

	if a.Config.ServerserviceOptions.OidcIssuerEndpoint == "" {
		return errors.New("serverservice oidc.issuer.endpoint not defined")
	}

	if a.v.GetString("serverservice.oidc.audience.endpoint") != "" {
		a.Config.ServerserviceOptions.OidcAudienceEndpoint = a.v.GetString("serverservice.oidc.audience.endpoint")
	}

	if a.Config.ServerserviceOptions.OidcAudienceEndpoint == "" {
		return errors.New("serverservice oidc.audience.endpoint not defined")
	}

	if a.v.GetString("serverservice.oidc.client.secret") != "" {
		a.Config.ServerserviceOptions.OidcClientSecret = a.v.GetString("serverservice.oidc.client.secret")
	}

	if a.Config.ServerserviceOptions.OidcClientSecret == "" {
		return errors.New("serverservice oidc.client.secret not defined")
	}

	if a.v.GetString("serverservice.oidc.client.id") != "" {
		a.Config.ServerserviceOptions.OidcClientID = a.v.GetString("serverservice.oidc.client.id")
	}

	if a.Config.ServerserviceOptions.OidcClientID == "" {
		return errors.New("serverservice oidc.client.id not defined")
	}

	if a.v.GetString("serverservice.oidc.client.scopes") != "" {
		a.Config.ServerserviceOptions.OidcClientScopes = a.v.GetStringSlice("serverservice.oidc.client.scopes")
	}

	if len(a.Config.ServerserviceOptions.OidcClientScopes) == 0 {
		return errors.New("serverservice oidc.client.scopes not defined")
	}

	return nil
}
