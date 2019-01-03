package omnicore

type Streamer interface {
        Cancel()
        OnCancel(func())
}

type Producer interface {
        Streamer

        Request(numElements uint32)
        OnData(callback func([]byte))
        OnEnd(callback func())

        Pipe(consumer Consumer)
}

type Consumer interface {
        Streamer

        Write(element []byte)
        End()
        OnRequest(func(numElements uint32))
        OnFinished(func())
}
