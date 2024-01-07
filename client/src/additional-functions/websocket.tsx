export class WebSocketService {
    private ws: WebSocket | null = null
    public isConnected = false

    private initializeSocket(url: string): void {
        this.ws = new WebSocket(url)

        this.ws.onopen = () => {
            this.isConnected = true
        }

        this.ws.onmessage = () => {

        }
    }
}