package token

import "time"

type Options struct {
	SecretKey []byte
}

type Option func(o *Options)

func WithSecretKey(key []byte) Option {
	return func(o *Options) {
		o.SecretKey = key
	}
}

func NewOptions(opts ...Option) Options {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return options
}

type GenerateOptions struct {
	Expiry time.Duration
}

type GenerateOption func(o *GenerateOptions)

func WithExpiry(d time.Duration) GenerateOption {
	return func(o *GenerateOptions) {
		o.Expiry = d
	}
}

func NewGenerateOptions(opts ...GenerateOption) GenerateOptions {
	var options GenerateOptions
	for _, o := range opts {
		o(&options)
	}
	//set default Expiry of token
	if options.Expiry == 0 {
		options.Expiry = time.Minute * 15
	}
	return options
}
