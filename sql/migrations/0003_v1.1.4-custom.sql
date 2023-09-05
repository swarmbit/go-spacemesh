ALTER TABLE rewards ADD atx_id CHAR(32) DEFAULT NULL;
CREATE INDEX rewards_by_atx_id ON rewards (atx_id, layer);
