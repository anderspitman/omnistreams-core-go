package omnicore

type Streamer interface {
        Cancel()
}

type Producer interface {
        Streamer
        OnData(callback func([]byte))
        OnEnd(callback func())
        Request(numElements uint8)
}

