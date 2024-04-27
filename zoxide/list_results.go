package zoxide

import (
	"strings"

	"github.com/joshmedeski/sesh/convert"
	"github.com/joshmedeski/sesh/model"
)

func (z *RealZoxide) ListResults() ([]model.ZoxideResult, error) {
	list, err := z.shell.ListCmd("zoxide", "query", "-l")
	if err != nil {
		return nil, err
	}
	results := make([]model.ZoxideResult, 0, len(list))
	for _, result := range list {
		trimmedResult := strings.TrimSpace(result)
		fields := strings.SplitN(trimmedResult, " ", 2)
		results = append(results, model.ZoxideResult{
			Score: convert.StringToFloat(fields[0]),
			Path:  fields[1],
		})
	}
	return results, nil
}
