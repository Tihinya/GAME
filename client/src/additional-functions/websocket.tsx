import { EventType, EventData } from "../types/Types";

export const ws = new WebSocket(`ws://${location.hostname}:8080/ws`);

const eventCallbacks = new Map<EventType, ((data: any) => void)[]>();

ws.onmessage = (event) => {
  const { type, payload }: EventData = JSON.parse(event.data);

  if (eventCallbacks.has(type)) {
    eventCallbacks.get(type)?.forEach((callback) => {
      callback(payload);
    });
  }
};

export const subscribe = (type: EventType, callback: (data: any) => void) => {
  if (!eventCallbacks.has(type)) {
    eventCallbacks.set(type, []);
  }
  eventCallbacks.get(type)?.push(callback);
};

export const unsubscribe = (type: EventType, callback: (data: any) => void) => {
  const callbacks = eventCallbacks.get(type);
  if (callbacks) {
    eventCallbacks.set(
      type,
      callbacks.filter((cb) => cb !== callback)
    );
  }
};

export const triggerEvent = (eventType: EventType, eventData: any) => {
  const callbacks = eventCallbacks.get(eventType);
  if (callbacks) {
    callbacks.forEach((callback) => {
      callback(eventData);
    });
  }
};
