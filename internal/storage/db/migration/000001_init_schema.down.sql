ALTER TABLE IF EXISTS students DROP CONSTRAINT IF EXISTS students_id_fkey;

ALTER TABLE IF EXISTS students DROP CONSTRAINT IF EXISTS students_group_id_fkey;

ALTER TABLE IF EXISTS mentors DROP CONSTRAINT IF EXISTS mentors_id_fkey;

ALTER TABLE IF EXISTS admins DROP CONSTRAINT IF EXISTS admins_id_fkey;

ALTER TABLE IF EXISTS student_details DROP CONSTRAINT IF EXISTS student_details_student_id_fkey;

ALTER TABLE IF EXISTS student_details DROP CONSTRAINT IF EXISTS student_details_reviewer_id_fkey;

ALTER TABLE IF EXISTS student_skills DROP CONSTRAINT IF EXISTS student_skills_student_id_fkey;

ALTER TABLE IF EXISTS student_skills DROP CONSTRAINT IF EXISTS student_skills_skill_fkey;

ALTER TABLE IF EXISTS contacts DROP CONSTRAINT IF EXISTS contacts_student_id_fkey;

ALTER TABLE IF EXISTS groups DROP CONSTRAINT IF EXISTS groups_mentor_fkey;

ALTER TABLE IF EXISTS group_overview DROP CONSTRAINT IF EXISTS group_overview_group_id_fkey;

ALTER TABLE IF EXISTS group_overview DROP CONSTRAINT IF EXISTS group_overview_reviewer_id_fkey;

ALTER TABLE IF EXISTS homework DROP CONSTRAINT IF EXISTS homework_student_id_fkey;

ALTER TABLE IF EXISTS diploma DROP CONSTRAINT IF EXISTS diploma_student_id_fkey;

ALTER TABLE IF EXISTS cert DROP CONSTRAINT IF EXISTS cert_student_id_fkey;

ALTER TABLE IF EXISTS interview DROP CONSTRAINT IF EXISTS interview_student_id_fkey;

ALTER TABLE IF EXISTS interview DROP CONSTRAINT IF EXISTS interview_mentor_id_fkey;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS students;

DROP TABLE IF EXISTS mentors;

DROP TABLE IF EXISTS admins;

DROP TABLE IF EXISTS student_details;

DROP TABLE IF EXISTS skills;

DROP TABLE IF EXISTS student_skills;

DROP TABLE IF EXISTS contacts;

DROP TABLE IF EXISTS groups;

DROP TABLE IF EXISTS group_overview;

DROP TABLE IF EXISTS homework;

DROP TABLE IF EXISTS diploma;

DROP TABLE IF EXISTS cert;

DROP TABLE IF EXISTS interview;
