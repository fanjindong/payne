package msg

import (
	"reflect"
	"testing"
)

func TestMapMsg(t *testing.T) {
	tests := []struct {
		name    string
		m       MapMsg
		want    MapMsg
		wantErr bool
	}{
		{name: "1", m: MapMsg{"a": 1}, want: MapMsg{"a": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			m := MapMsg{}
			err = m.Unmarshal(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(m, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", m, tt.want)
			}
		})
	}
}
