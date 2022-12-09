package main

import (
	"reflect"
	"testing"
)

const testMultiLine = `
Running Servers Statistics
===========================
Children number: 3
Free Servers: 28
Used Servers: 2
Started Processes: 10
Closed Processes: 7
Crashed Processes: 0
Closing Processes: 0

Child pids: 19 73 85
Closing children pids:
Semaphores in use
     file:/tmp/icap_lock_accept.Z1vVqw
     file:/tmp/icap_lock_children-queue.g6KY1v

Shared mem blocks in use
	 posix:/c-icap-shared-kids-queue.0 13 kbs


General Statistics
==================
REQUESTS : 12341
REQMODS : 0
RESPMODS : 12311
OPTIONS : 30
FAILED REQUESTS : 11
ALLOW 204 : 6320
BYTES IN : 5933142 Kbs 543 bytes
BYTES OUT : 1749866 Kbs 571 bytes
HTTP BYTES IN : 5925503 Kbs 102 bytes
HTTP BYTES OUT : 1745166 Kbs 477 bytes
BODY BYTES IN : 5909975 Kbs 646 bytes
BODY BYTES OUT : 1742211 Kbs 803 bytes

Service info Statistics
==================
Service info REQMODS : 0
Service info RESPMODS : 0
Service info OPTIONS : 1
Service info ALLOW 204 : 0
Service info BYTES IN : 0 Kbs 138 bytes
Service info BYTES OUT : 0 Kbs 297 bytes
Service info HTTP BYTES IN : 0 Kbs 0 bytes
Service info HTTP BYTES OUT : 0 Kbs 0 bytes
Service info BODY BYTES IN : 0 Kbs 0 bytes
Service info BODY BYTES OUT : 0 Kbs 0 bytes

Service echo Statistics
==================
Service echo REQMODS : 0
Service echo RESPMODS : 0
Service echo OPTIONS : 0
Service echo ALLOW 204 : 0
Service echo BYTES IN : 0 Kbs 0 bytes
Service echo BYTES OUT : 0 Kbs 0 bytes
Service echo HTTP BYTES IN : 0 Kbs 0 bytes
Service echo HTTP BYTES OUT : 0 Kbs 0 bytes
Service echo BODY BYTES IN : 0 Kbs 0 bytes
Service echo BODY BYTES OUT : 0 Kbs 0 bytes

Service virus_scan Statistics
==================
Service virus_scan REQMODS : 0
Service virus_scan RESPMODS : 12311
Service virus_scan OPTIONS : 29
Service virus_scan ALLOW 204 : 6320
Requests scanned : 9522
Virmode requests : 0
Viruses found : 10
Scan failures : 805
Service virus_scan BYTES IN : 5933142 Kbs 405 bytes
Service virus_scan BYTES OUT : 1749866 Kbs 274 bytes
Service virus_scan HTTP BYTES IN : 5925503 Kbs 102 bytes
Service virus_scan HTTP BYTES OUT : 1745166 Kbs 477 bytes
Service virus_scan BODY BYTES IN : 5909975 Kbs 646 bytes
Service virus_scan BODY BYTES OUT : 1742211 Kbs 803 bytes
Body bytes scanned : 5899516 Kbs 501 bytes
`

func Test_parseResult(t *testing.T) {
	type args struct {
		result string
	}
	tests := []struct {
		name    string
		args    args
		want    *Stats
		wantErr bool
	}{
		{
			name: "test multiline OK",
			args: args{result: testMultiLine},
			want: &Stats{
				generalRequestsTotal:           12341,
				generalFailedRequestsTotal:     11,
				virusScanRequestsRespModsTotal: 12311,
				virusScanRequestsScannedTotal:  9522,
				virusScanVirusFoundTotal:       10,
				virusScanFailureTotal:          805,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseResult(tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}
