// Package plugin is a starter togo plugin. Copy this repository (or use
// `togo make:plugin <name>`) and implement the lifecycle methods below.
//
// A plugin satisfies the togo.Plugin contract:
//
//	Name() string
//	Priority() int
//	Register(k *togo.Kernel) error
//	Boot(ctx context.Context, k *togo.Kernel) error
//
// On `togo install <owner>/<repo>`, the manifest (togo.plugin.yaml) is read and
// this package is registered with the kernel for auto-discovery. The togo import
// is added once you wire the methods to the kernel API.
package plugin

import "context"

// Plugin is the example plugin. Rename it for your plugin.
type Plugin struct{}

// Name uniquely identifies the plugin.
func (Plugin) Name() string { return "example" }

// Priority controls boot order (0–100, lower boots first).
func (Plugin) Priority() int { return 50 }

// Register binds services, config and hooks. Replace `any` with *togo.Kernel
// once you depend on github.com/togo-framework/togo.
func (Plugin) Register(k any) error { return nil }

// Boot mounts routes, registers schema, and runs migrations.
func (Plugin) Boot(ctx context.Context, k any) error { return nil }
