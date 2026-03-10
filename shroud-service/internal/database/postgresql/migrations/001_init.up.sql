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

CREATE TABLE interactions ();

CREATE TABLE users ();

CREATE TABLE credentials ();

CREATE TABLE profiles ();

CREATE TABLE devices ();

CREATE TABLE relationships ();

CREATE TABLE roles ();

CREATE TABLE members ();

CREATE TABLE member_roles ();

CREATE TABLE overrides ();

CREATE TABLE entitlements ();

CREATE TABLE applications ();

CREATE TABLE tiers ();