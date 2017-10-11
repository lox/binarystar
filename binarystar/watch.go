// +build !darwin

package binarystar

import (
	"context"

	"github.com/pkg/errors"
)

func watch(ctx context.Context, tree *Tree) error {
	return errors.New("Not implemented")
}
