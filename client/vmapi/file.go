/* SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) Benedict Schlueter
 */

package vmapi

import (
	"context"
	"os"
	"path/filepath"

	"github.com/benschlueter/delegatio/client/vmapi/vmproto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WriteFile creates a file and writes output to it.
func (a *API) WriteFile(ctx context.Context, in *vmproto.WriteFileRequest) (*vmproto.WriteFileResponse, error) {
	a.logger.Info("request to write file", zap.String("path", in.Filepath), zap.String("name", in.Filename))
	if err := os.WriteFile(filepath.Join(in.Filepath, in.Filename), in.Content, os.ModeAppend); err != nil {
		a.logger.Error("failed to write file", zap.String("path", in.Filepath), zap.String("name", in.Filename), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "file write failed exited with error code: %v", err)
	}
	return &vmproto.WriteFileResponse{}, nil
}
