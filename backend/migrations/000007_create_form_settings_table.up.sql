CREATE TABLE
    IF NOT EXISTS core.form_settings (
        id smallserial PRIMARY KEY,
        background_color char(7) NOT NULL,
        foreground_color char(7) NOT NULL,
        primary_color char(7) NOT NULL,
        created_at timestamptz NOT NULL DEFAULT now (),
        CONSTRAINT chk_background_color CHECK (background_color ~ '^#[0-9a-fA-F]{6}$'),
        CONSTRAINT chk_foreground_color CHECK (foreground_color ~ '^#[0-9a-fA-F]{6}$'),
        CONSTRAINT chk_primary_color CHECK (primary_color ~ '^#[0-9a-fA-F]{6}$')
    );
INSERT INTO
    core.form_settings (background_color, foreground_color, primary_color)
VALUES
    ('#FFFFFF', '#000000', '#FF5733'),
    ('#000000', '#FFFFFF', '#3498DB'),
    ('#F0F0F0', '#202020', '#2ECC71'),
    ('#282C34', '#ABB2BF', '#E06C75');