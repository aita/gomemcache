// Copyright 2018 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memcache

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

const (
	keyMethod = "method"
	keyStatus = "status"
)

type methodSpan struct {
	startTime  time.Time
	methodName string
	span       opentracing.Span
}

func newMethodSpan(ctx context.Context, methodName string, keys ...string) (*methodSpan, context.Context) {
	span := methodSpan{}
	ctx = span.start(ctx, methodName, keys...)
	return &span, ctx
}

func (span *methodSpan) start(ctx context.Context, methodName string, keys ...string) context.Context {
	span.startTime = time.Now()
	span.methodName = methodName
	span.span, ctx = opentracing.StartSpanFromContext(ctx, methodName,
		opentracing.Tag{keyMethod, span.methodName},
		opentracing.StartTime(span.startTime),
	)
	ext.DBType.Set(span.span, "memcached")
	return ctx
}

func (span *methodSpan) finish(err error) {
	if err == nil {
		span.span.SetTag(keyStatus, "OK")
	} else {
		span.span.SetTag(keyStatus, "ERROR")
		ext.Error.Set(span.span, true)
		span.span.LogFields(
			log.String("event", "error"),
			log.String("message", err.Error()),
		)
	}
	span.span.Finish()
}
