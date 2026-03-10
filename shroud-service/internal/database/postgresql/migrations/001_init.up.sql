CREATE TABLE realms (
    id uuid PRIMARY KEY,
    owner_id uuid,
    creation_time timestamptz,
    first_seen timestamptz,
    last_seen timestamptz,
    allows_nsfw bool
);

CREATE TABLE guilds (
    id uuid PRIMARY KEY,
    owner_id uuid,
    creation_time timestamptz,
    nsfw bool,
    name varchar(25),
    description varchar(200),
    icon varchar,
    banner varchar,
    invite_background varchar,
    tags varchar(15)[],
    welcome_message varchar(1000),
    support_message varchar(1000),
    user_feed bool,
    notifications_all bool,
    realm_general_notification_channel uuid,
    realm_moderation_notification_channel uuid,
    access_type varchar,
    application_questions jsonb,
    suspicious_activity_notifications uuid,
    verification_level varchar,
    allow_dms_from_suspicious bool,
    allow_dms_from_unknown bool,
    require_moderator_2fa bool,
    automod_rules jsonb
);

CREATE TABLE channels (
    id uuid PRIMARY KEY,
    parent_id uuid,
    guild_id uuid,
    creation_time timestamptz,
    name varchar(25),
    type varchar,
    nsfw bool,
    slowmode int,
    topic varchar(500),
    bitrate int,
    user_limit int,
    tags jsonb,
    default_layout varchar,
    default_order varchar,
    locked bool,
    private bool,
    image varchar
);

CREATE TABLE interactions (
    id uuid PRIMARY KEY,
    type varchar,
    name varchar(20),
    icon uuid,
    asset varchar
);

CREATE TABLE users ();

CREATE TABLE user_credentials (
    id uuid
);

CREATE TABLE user_profiles ();

CREATE TABLE user_devices ();

CREATE TABLE user_relationships ();

CREATE TABLE guild_invites (
    id uuid,
    code varchar,
    generated_by uuid,
    generate_at timestamptz,
    expires_at timestamptz,
    use_limit int,
    bypass_application bool
);

CREATE TABLE guild_roles (
    id uuid PRIMARY KEY,
    guild_id uuid,
    name varchar(25),
    color varchar(6),
    position int,
    permissions bigint
);

CREATE TABLE guild_members (
    guild_id uuid,
    user_id uuid,
    PRIMARY KEY (guild_id,user_id),

    role_ids uuid[],

    application_status varchar,
    joined_at timestamptz,
    join_method varchar,
    join_link varchar,
    join_link_generated_by uuid,

    guild_profile uuid
);

CREATE TABLE member_roles (
    user_id uuid,
    guild_id uuid,
    role_id uuid,
    PRIMARY KEY (user_id, guild_id)
);

CREATE TABLE channel_overrides (
    guild_id uuid,
    channel_id uuid,
    PRIMARY KEY (guild_id,channel_id),

    scope varchar,
    object_id uuid,
    permissions_allow bigint
);

CREATE TABLE applications ();

CREATE TABLE tiers (
    id uuid PRIMARY KEY,
    name varchar(50),
    badge varchar,
    tier_expiry timestamptz,
    upload_size int
);