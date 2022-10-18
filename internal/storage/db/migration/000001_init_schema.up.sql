CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "birthday" datetime NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL
);

CREATE TABLE "students" (
  "id" int NOT NULL,
  "group_id" varchar NOT NULL
);

CREATE TABLE "mentors" (
  "id" int
);

CREATE TABLE "admins" (
  "id" int
);

CREATE TABLE "student_details" (
  "student_id" int,
  "resume" text,
  "mentor_review" text,
  "reviewer_id" int
);

CREATE TABLE "skills" (
  "skill" varchar PRIMARY KEY
);

CREATE TABLE "student_skills" (
  "student_id" int NOT NULL,
  "skill" int NOT NULL,
  "level" int NOT NULL
);

CREATE TABLE "contacts" (
  "student_id" int,
  "discord" varchar,
  "telegram" varchar,
  "instagram" varchar
);

CREATE TABLE "groups" (
  "id" varchar UNIQUE PRIMARY KEY,
  "mentor" int NOT NULL,
  "date_start" datetime NOT NULL,
  "date_end" datetime NOT NULL
);

CREATE TABLE "group_overview" (
  "group_id" int,
  "overview" text NOT NULL,
  "reviewer_id" int NOT NULL
);

CREATE TABLE "homework" (
  "student_id" int,
  "hommework_score_1" int,
  "hommework_score_2" int,
  "hommework_score_3" int,
  "hommework_score_4" int,
  "hommework_score_5" int,
  "hommework_score_6" int
);

CREATE TABLE "diploma" (
  "student_id" int NOT NULL,
  "theme" varchar NOT NULL,
  "description" text
);

CREATE TABLE "cert" (
  "student_id" id NOT NULL,
  "cert_id" varchar UNIQUE NOT NULL,
  "issued" datetime NOT NULL
);

CREATE TABLE "interview" (
  "student_id" int,
  "mentor_id" int,
  "date" datetime NOT NULL,
  "mark" int NOT NULL,
  "notes" text
);

ALTER TABLE "students" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "students" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "mentors" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "admins" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "student_details" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_details" ADD FOREIGN KEY ("reviewer_id") REFERENCES "mentors" ("id");

ALTER TABLE "student_skills" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "student_skills" ADD FOREIGN KEY ("skill") REFERENCES "skills" ("skill");

ALTER TABLE "contacts" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "groups" ADD FOREIGN KEY ("mentor") REFERENCES "mentors" ("id");

ALTER TABLE "group_overview" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "group_overview" ADD FOREIGN KEY ("reviewer_id") REFERENCES "mentors" ("id");

ALTER TABLE "homework" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "diploma" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "cert" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "interview" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "interview" ADD FOREIGN KEY ("mentor_id") REFERENCES "mentors" ("id");
