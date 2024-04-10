CREATE TABLE Account(
	sum float(32) NOT NULL CHECK(sum >= 0.0)
);

INSERT INTO Account VALUES (1000.0);