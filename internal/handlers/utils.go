package handlers

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseId(data string) (uint, error) {
	params := strings.Split(data, " ")
	if len(params) != 1 {
		return 0, errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params)
	}
	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return 0, errors.Wrapf(BadArgument, "Invalid id: %s", params[0])
	}
	return uint(id), nil
}
