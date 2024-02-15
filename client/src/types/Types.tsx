export enum EventType {
  EventSendMessage = "send_message",
  EventAmaBoy = "ama_boy_next_door",
  GameEventNotification = "game_notification",
  GameEventMovePlayer = "game_move",
  EventLoginHandler = "register_user",
  EventOnlineUserList = "online_users_list",
  EventReceiveMessage = "receive_message",
  EventClientInfoMessage = "client_info",
  GameEventError = "game_error",
  GameEventInput = "game_input",
  GameEventGameState = "game_state",
  GameEventBomb = "game_bomb",
  GameEventObstacle = "game_obstacle",
  GameEventPowerup = "game_powerup",
  GameEvent = "game_event",
  GameEventPlayerMotion = "game_player_position",
  GameEventPlayerHealth = "game_player_health",
  GameEventPlayerCreation = "game_player_creation",
}

export type EventData<T = any> = {
  type: EventType;
  payload: T;
};

export type Players = {
  [key: string]: null;
};

export type ChatMessage = {
  name: string;
  message: string;
};

export type TimeCountdown = {
  state: string;
  currentTime: number;
};

export enum PageState {
  MainPage = "main_page",
  Lobby = "lobby",
  GamePage = "game_page",
}

export type Page = {
  state: PageState;
};

type Position = {
  x: number;
  y: number;
  id: number;
  name: string;
};

type Block = {
  name: string;
};

export type GameState = {
  players: Position[];
  powerups: Position[];
  map: (Block | null)[][];
};
