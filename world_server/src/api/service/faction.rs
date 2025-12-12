pub fn get_faction_id(name: String) -> Option<u8> {
    match name.to_lowercase().as_str() {
        "caldari" => Some(1),
        "varnak" => Some(2),
        "dawnhold" => Some(3),
        _ => None,
    }
}