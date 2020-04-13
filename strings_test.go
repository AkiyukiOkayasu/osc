/*
Copyright 2020 Akiyuki Okayasu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package osc

import (
	"reflect"
	"testing"
)

func Test_splitOSCPacket(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		wantM Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := splitOSCPacket(tt.args.str); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("splitOSCPacket() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_numNeededNullChar(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 3},
		{"2", args{2}, 2},
		{"3", args{3}, 1},
		{"4", args{4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numNeededNullChar(tt.args.l); got != tt.want {
				t.Errorf("numNeededNullChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split2OSCStrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{",i", args{",i" + string(nullChar) + string(nullChar) + string('\x00') + string('\x00') + string('\x00') + string('\x01')}, ",i", string('\x00') + string('\x00') + string('\x00') + string('\x01')},
		{",ff", args{",ff" + string(nullChar) + string('\x00') + string('\x00') + string('\x00') + string('\x01') + string('\x00') + string('\x00') + string('\x00') + string('\x01')}, ",ff", string('\x00') + string('\x00') + string('\x00') + string('\x01') + string('\x00') + string('\x00') + string('\x00') + string('\x01')},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := split2OSCStrings(tt.args.s)
			if got != tt.want {
				t.Errorf("split2OSCStrings() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("split2OSCStrings() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_terminateOSCString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{",i", args{",i"}, ",i" + string(nullChar) + string(nullChar)},
		{",if", args{",if"}, ",if" + string(nullChar)},
		{",ifs", args{",ifs"}, ",ifs" + string(nullChar) + string(nullChar) + string(nullChar) + string(nullChar)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := terminateOSCString(tt.args.str); got != tt.want {
				t.Errorf("terminateOSCString() = %v, want %v", got, tt.want)
			}
		})
	}
}
