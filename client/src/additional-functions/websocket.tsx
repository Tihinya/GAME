export const ws = new WebSocket(`ws://${location.hostname}:8080/ws`);

const eventCallbacks = new Map<any, ((data: any) => void)[]>();

ws.onmessage = (event) => {
  const { type, payload }: any = JSON.parse(event.data);
  if (eventCallbacks.has(type)) {
    eventCallbacks.get(type)?.forEach((callback) => {
      callback(payload);
    });
  }
};

export const subscribe = (type: any, callback: (data: any) => void) => {
  if (!eventCallbacks.has(type)) {
    eventCallbacks.set(type, []);
  }
  eventCallbacks.get(type)?.push(callback);
};

export const unsubscribe = (type: any, callback: (data: any) => void) => {
  const callbacks = eventCallbacks.get(type);
  if (callbacks) {
    eventCallbacks.set(
      type,
      callbacks.filter((cb) => cb !== callback)
    );
  }
};

// export const triggerEvent = (type: any, eventData: any) => {
//   const callbacks = eventCallbacks.get(type);
//   if (callbacks) {
//     callbacks.forEach((callback) => {
//       callback(eventData);
//     });
//   }
// };
