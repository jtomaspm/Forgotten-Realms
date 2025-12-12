use api_sdk::{requests, responses};
use axum::{Extension, Json, Router, routing::post};
use tokio::sync::mpsc::Sender;
use uuid::Uuid;

use crate::{api::{models::ErrorResponse, service::faction::{get_faction_id}}, event_hub::event::{Event, EventPayload, player::CreatePlayerEvent}};


pub fn mount(router: Router) -> Router {
    return router.route("/player", post(create_player));
}

async fn create_player(hub_channel: Extension<Sender<Event>>, Json(request): Json<requests::realm_server::CreatePlayer>) -> Result<Json<responses::realm_server::CreatePlayer>, ErrorResponse> {
    match get_faction_id(request.faction.clone()) {
        Some(faction_id) => {
            hub_channel.send(Event {
                id: Uuid::new_v4(),
                response_channel: None,
                payload: EventPayload::CreatePlayer(CreatePlayerEvent {
                    name: request.name.clone(),
                    faction_id: faction_id,
                })
            }).await.unwrap();

            return Ok(Json(responses::realm_server::CreatePlayer {
                player_id: 1,
            }));
        },
        None => {
            return Err(ErrorResponse {
                error: format!("Invalid faction name: {}", request.faction),
            });
        },
    }
}