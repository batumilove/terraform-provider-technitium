// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/darkhonor/terraform-provider-technitium/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &ServerSettingsDataSource{}

func NewServerSettingsDataSource() datasource.DataSource {
	return &ServerSettingsDataSource{}
}

type ServerSettingsDataSource struct {
	client *client.Client
}

type ServerSettingsDataSourceModel struct {
	ID                          types.String `tfsdk:"id"`
	Version                     types.String `tfsdk:"version"`
	Uptime                      types.String `tfsdk:"uptime"`
	DnssecValidation            types.Bool   `tfsdk:"dnssec_validation"`
	Recursion                   types.String `tfsdk:"recursion"`
	QnameMinimization           types.Bool   `tfsdk:"qname_minimization"`
	RandomizeName               types.Bool   `tfsdk:"randomize_name"`
	LogQueries                  types.Bool   `tfsdk:"log_queries"`
	LoggingType                 types.String `tfsdk:"logging_type"`
	MaxLogFileDays              types.Int64  `tfsdk:"max_log_file_days"`
	EnableBlocking              types.Bool   `tfsdk:"enable_blocking"`
	ServeStale                  types.Bool   `tfsdk:"serve_stale"`
	ForwarderProtocol           types.String `tfsdk:"forwarder_protocol"`
	EnableDnsOverTls            types.Bool   `tfsdk:"enable_dns_over_tls"`
	EnableDnsOverHttps          types.Bool   `tfsdk:"enable_dns_over_https"`
	UdpPayloadSize              types.Int64  `tfsdk:"udp_payload_size"`
	CacheMinimumRecordTtl       types.Int64  `tfsdk:"cache_minimum_record_ttl"`
	CacheMaximumRecordTtl       types.Int64  `tfsdk:"cache_maximum_record_ttl"`
}

func (d *ServerSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server_settings"
}

func (d *ServerSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads current Technitium DNS Server settings for compliance auditing.",
		Attributes: map[string]schema.Attribute{
			"id":                     schema.StringAttribute{Computed: true, Description: "Fixed identifier."},
			"version":                schema.StringAttribute{Computed: true, Description: "Server version."},
			"uptime":                 schema.StringAttribute{Computed: true, Description: "Server uptime timestamp."},
			"dnssec_validation":      schema.BoolAttribute{Computed: true, Description: "DNSSEC validation enabled."},
			"recursion":              schema.StringAttribute{Computed: true, Description: "Recursion policy."},
			"qname_minimization":     schema.BoolAttribute{Computed: true, Description: "QNAME minimization enabled."},
			"randomize_name":         schema.BoolAttribute{Computed: true, Description: "Query name randomization enabled."},
			"log_queries":            schema.BoolAttribute{Computed: true, Description: "Query logging enabled."},
			"logging_type":           schema.StringAttribute{Computed: true, Description: "Logging output type."},
			"max_log_file_days":      schema.Int64Attribute{Computed: true, Description: "Max log retention days."},
			"enable_blocking":        schema.BoolAttribute{Computed: true, Description: "DNS blocking enabled."},
			"serve_stale":            schema.BoolAttribute{Computed: true, Description: "Serve stale records enabled."},
			"forwarder_protocol":     schema.StringAttribute{Computed: true, Description: "Forwarder protocol."},
			"enable_dns_over_tls":    schema.BoolAttribute{Computed: true, Description: "DNS-over-TLS enabled."},
			"enable_dns_over_https":  schema.BoolAttribute{Computed: true, Description: "DNS-over-HTTPS enabled."},
			"udp_payload_size":       schema.Int64Attribute{Computed: true, Description: "EDNS UDP payload size."},
			"cache_minimum_record_ttl": schema.Int64Attribute{Computed: true, Description: "Minimum cache TTL."},
			"cache_maximum_record_ttl": schema.Int64Attribute{Computed: true, Description: "Maximum cache TTL."},
		},
	}
}

func (d *ServerSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	providerData, ok := req.ProviderData.(*TechnitiumProviderData)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *TechnitiumProviderData, got: %T", req.ProviderData))
		return
	}
	d.client = providerData.Client
}

func (d *ServerSettingsDataSource) Read(_ context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	settings, err := d.client.SettingsGet()
	if err != nil {
		resp.Diagnostics.AddError("Error reading server settings", err.Error())
		return
	}

	state := ServerSettingsDataSourceModel{
		ID:                    types.StringValue("server-settings"),
		Version:               types.StringValue(settings.Version),
		Uptime:                types.StringValue(settings.Uptimestamp),
		DnssecValidation:      types.BoolValue(settings.DnssecValidation),
		Recursion:             types.StringValue(settings.Recursion),
		QnameMinimization:     types.BoolValue(settings.QnameMinimization),
		RandomizeName:         types.BoolValue(settings.RandomizeName),
		LogQueries:            types.BoolValue(settings.LogQueries),
		LoggingType:           types.StringValue(settings.LoggingType),
		MaxLogFileDays:        types.Int64Value(int64(settings.MaxLogFileDays)),
		EnableBlocking:        types.BoolValue(settings.EnableBlocking),
		ServeStale:            types.BoolValue(settings.ServeStale),
		ForwarderProtocol:     types.StringValue(settings.ForwarderProtocol),
		EnableDnsOverTls:      types.BoolValue(settings.EnableDnsOverTls),
		EnableDnsOverHttps:    types.BoolValue(settings.EnableDnsOverHttps),
		UdpPayloadSize:        types.Int64Value(int64(settings.UdpPayloadSize)),
		CacheMinimumRecordTtl: types.Int64Value(int64(settings.CacheMinimumRecordTtl)),
		CacheMaximumRecordTtl: types.Int64Value(int64(settings.CacheMaximumRecordTtl)),
	}

	resp.Diagnostics.Append(resp.State.Set(context.Background(), &state)...)
}
