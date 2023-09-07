package optional

import (
	"encoding/json"
	"reflect"
	"testing"
)

var (
	fakeOptional = Optional[string]{
		Set:    true,
		IsNull: false,
		Value:  "1234",
	}
	fakeByteOptional, _     = json.Marshal(&fakeOptional)
	fakeByteOptionalNull, _ = json.Marshal(&Optional[string]{
		Set:    false,
		IsNull: true,
		Value:  "",
	})
)

func TestOptional_HasValue(t *testing.T) {
	type fields struct {
		Optional Optional[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success, has value true",
			fields: fields{
				Optional: Optional[string]{
					Set:    true,
					IsNull: false,
					Value:  "1234",
				},
			},
			want: true,
		},
		{
			name: "success, has value false",
			fields: fields{
				Optional: Optional[string]{
					Set:    false,
					IsNull: true,
					Value:  "1234",
				},
			},
			want: false,
		},
		{
			name: "success, has value false",
			fields: fields{
				Optional: Optional[string]{
					Set:    true,
					IsNull: true,
					Value:  "1234",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Optional[string]{
				Set:    tt.fields.Optional.Set,
				IsNull: tt.fields.Optional.IsNull,
				Value:  tt.fields.Optional.Value,
			}
			got := o.HasValue()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_GetValuePointer(t *testing.T) {
	type fields struct {
		Optional Optional[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "success, has value true",
			fields: fields{
				Optional: fakeOptional,
			},
			want: pointer(fakeOptional),
		},
		{
			name: "success, has value false",
			fields: fields{
				Optional: Optional[string]{
					Set:    false,
					IsNull: true,
					Value:  "1234",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Optional[string]{
				Set:    tt.fields.Optional.Set,
				IsNull: tt.fields.Optional.IsNull,
				Value:  tt.fields.Optional.Value,
			}
			got := o.GetValuePointer()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValuePointer()) = %v, want %v", got, tt.want)
			}
		})
	}
}

func pointer(value Optional[string]) *string {
	return &value.Value
}

func TestOptional_MarshalJSON(t *testing.T) {
	type fields struct {
		Optional Optional[string]
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "success, marshal success",
			fields: fields{
				Optional: fakeOptional,
			},
			want:    fakeByteOptional,
			wantErr: false,
		},
		{
			name: "success, marshal success",
			fields: fields{
				Optional: Optional[string]{
					Set:    false,
					IsNull: true,
					Value:  "",
				},
			},
			want:    fakeByteOptionalNull,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Optional[string]{
				Set:    tt.fields.Optional.Set,
				IsNull: tt.fields.Optional.IsNull,
				Value:  tt.fields.Optional.Value,
			}
			got, err := o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
