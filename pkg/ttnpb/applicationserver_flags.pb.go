// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.6
// - protoc              v3.21.1
// source: lorawan-stack/api/applicationserver.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// AddSelectFlagsForApplicationLink adds flags to select fields in ApplicationLink.
func AddSelectFlagsForApplicationLink(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("default-formatters", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("default-formatters", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForMessagePayloadFormatters(flags, flagsplugin.Prefix("default-formatters", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("skip-payload-crypto", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("skip-payload-crypto", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forApplicationLink message from select flags.
func PathsFromSelectFlagsForApplicationLink(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("default_formatters", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("default_formatters", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForMessagePayloadFormatters(flags, flagsplugin.Prefix("default_formatters", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("skip_payload_crypto", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("skip_payload_crypto", prefix))
	}
	return paths, nil
}

// AddSetFlagsForApplicationLink adds flags to select fields in ApplicationLink.
func AddSetFlagsForApplicationLink(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForMessagePayloadFormatters(flags, flagsplugin.Prefix("default-formatters", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("skip-payload-crypto", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the ApplicationLink message from flags.
func (m *ApplicationLink) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("default_formatters", prefix)); changed {
		if m.DefaultFormatters == nil {
			m.DefaultFormatters = &MessagePayloadFormatters{}
		}
		if setPaths, err := m.DefaultFormatters.SetFromFlags(flags, flagsplugin.Prefix("default_formatters", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("skip_payload_crypto", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.SkipPayloadCrypto = &wrapperspb.BoolValue{Value: val}
		paths = append(paths, flagsplugin.Prefix("skip_payload_crypto", prefix))
	}
	return paths, nil
}
