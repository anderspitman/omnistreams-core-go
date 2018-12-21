package omnicore

type Producer interface {
        OnData(callback func([]byte))
        OnEnd(callback func())
        Request(numElements uint8)
}

