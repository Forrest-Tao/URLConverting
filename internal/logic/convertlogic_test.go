package logic

import (
	"context"
	"reflect"
	"shorturl/internal/svc"
	"shorturl/internal/types"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestConvertLogic_Convert(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.ConvertRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.ConvertResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ConvertLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			gotResp, err := l.Convert(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Convert() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestNewConvertLogic(t *testing.T) {
	type args struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	tests := []struct {
		name string
		args args
		want *ConvertLogic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConvertLogic(tt.args.ctx, tt.args.svcCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConvertLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}
