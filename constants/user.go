package constants

import "time"

// topics
const Jetson_ANPR string = "_jetson"
const Parkzap_epc string = "_parkbox"
const Parkzap_loop string = "_localfirst"

// msg codes
const Parkzap_Msg_Code float64 = 8000
const Jetson_Msg_Code float64 = 10002
const Parkzap_loop_code float64 = 8027

// channel buffer(number of values queued)
const ChannelBuffer float32 = 1

// qos(quality of service)
const Qos float32 = 0

// redis db
const Redis_db int = 0

// Broker Protocol
const Broker_Protocol string = "tcp://"

// number of channels
const No_channels int = 1

// redis timeout time
const Redis_time_out time.Duration = 5 * time.Second

// keys
const Epc_key string = "Epc"
const Anpr_key string = "ANPR1"
const Loop_key string = "LOOP"

// The sample rate for sampling traces in the range [0.0, 1.0].
const TracesSampleRate float64 = 1.0

// length of Epc
const LenEpc int = 24
