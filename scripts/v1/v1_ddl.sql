drop table if exists  reports;
drop table if exists  parameters;
drop table if exists  parameter_groups;
drop table if exists  parameter_rules;
drop table if exists  result_type;

-- registrations
drop table if exists  registrations;
drop table if exists  work_orders;
drop table if exists  users;
drop table if exists  campaigns;
drop table if exists  pin_codes;

drop table if exists  roles;
drop table if exists  departments;
drop table if exists  stores;
drop table if exists  assigning_authority;
drop table if exists  application_access;
drop table if exists  districts;
drop table if exists  states;

-- roles
CREATE TABLE IF NOT EXISTS roles (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(18) NOT NULL
);


CREATE TABLE IF NOT EXISTS departments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(18) NOT NULL
);


-- campaign

CREATE TABLE IF NOT EXISTS stores (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(44) NOT NULL
);

CREATE TABLE IF NOT EXISTS assigning_authority (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(18) NOT NULL
);

CREATE TABLE IF NOT EXISTS application_access (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(18) NOT NULL
);

CREATE TABLE IF NOT EXISTS states (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS districts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(40) NOT NULL,
  state_id UUID NOT NULL,
  FOREIGN KEY(state_id) REFERENCES states(id)
);

CREATE TABLE IF NOT EXISTS pin_codes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  pin_code VARCHAR(40) NOT NULL,
  district_id UUID NOT NULL,
  FOREIGN KEY(district_id) REFERENCES districts(id)
);

CREATE TABLE IF NOT EXISTS campaigns (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  distict VARCHAR(100) NOT NULL,
  village VARCHAR(100) NOT NULL,
  taluk_name VARCHAR(100) NOT NULL,
  pin_code VARCHAR(100) NOT NULL,
  camp_id VARCHAR(100) NOT NULL,
  work_order VARCHAR(100) NOT NULL,
  visibility VARCHAR(100) NOT NULL,
  status VARCHAR(100) NOT NULL,
  created_by VARCHAR(100) NOT NULL,
  created_at Date NOT NULL,
  updated_at Date NOT NULL,
  state_name VARCHAR(100) ,
  description TEXT NOT NULL,

  estimated_target_screening INTEGER NOT NULL,
  labour_inspector_name VARCHAR(100) ,
  union_name VARCHAR(100) ,
  union_leader_name VARCHAR(100) ,
  latitude VARCHAR(100) NOT NULL,
  longitude VARCHAR(100) NOT NULL,
  
  screening_start_date DATE NOT NULL,
  screening_start_time TEXT NOT NULL,

  application_access_id UUID ,
  assigning_authority_id UUID ,
  store_id UUID ,

  FOREIGN KEY(application_access_id) REFERENCES application_access(id),
  FOREIGN KEY(assigning_authority_id) REFERENCES assigning_authority(id),
  FOREIGN KEY(store_id) REFERENCES stores(id)
);

-- user
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  user_name VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(18) NOT NULL,
  mobile VARCHAR(11) NOT NULL,
  email_id VARCHAR(100) NOT NULL,
  assigning_authority_id UUID ,
  role_id UUID NOT NULL,
  department_id UUID ,
  application_access_id UUID ,
  gender VARCHAR(1) NOT NULL,
  first_name VARCHAR(18) NOT NULL UNIQUE,
  last_name VARCHAR(18) NOT NULL UNIQUE,
  middle_name VARCHAR(18) NOT NULL UNIQUE,
  date_of_birth DATE NOT NULL,
  date_of_joining DATE NoT NULL,
  last_working_day DATE ,  
  monthly_ctc VARCHAR(100) NoT NULL,
  campaign_id UUID,
  
  FOREIGN KEY(assigning_authority_id) REFERENCES assigning_authority(id),
  FOREIGN KEY(department_id) REFERENCES departments(id),
  FOREIGN KEY(application_access_id) REFERENCES application_access(id),
  FOREIGN KEY(campaign_id) REFERENCES campaigns(id),
  FOREIGN KEY(role_id) REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS work_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(40) NOT NULL
);
-- patient registration
CREATE TABLE IF NOT EXISTS registrations (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  registration_date DATE NOT NULL,
  uhid VARCHAR(40) NOT NULL,
  barcode VARCHAR(40) NOT NULL,
  name VARCHAR(40) NOT NULL,
  labour_id VARCHAR(40) NOT NULL,
  age INTEGER NOT NULL,
  gender VARCHAR(1) NOT NULL,
  mobile VARCHAR(11) NOT NULL,
  taluk VARCHAR(40) NOT NULL,
  lab_test_status INTEGER,
  report_url TEXT,

  campaign_id UUID NOT NULL,
  district_id UUID NOT NULL,
  FOREIGN KEY(campaign_id) REFERENCES campaigns(id),
  FOREIGN KEY(district_id) REFERENCES districts(id)
);

-- patient report

CREATE TABLE IF NOT EXISTS result_type (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  type INTEGER NOT NULL,
  drop_down_table_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS parameter_rules (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	expression TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS parameter_groups (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	group_name VARCHAR(40) NOT NULL,
	group_sequence INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS parameters (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  parameter_name VARCHAR(40) NOT NULL,
  result_type_id UUID NOT NULL,
  result VARCHAR(40) NOT NULL,
  unit VARCHAR(40) NOT NULL,
  interpretation VARCHAR(40) NOT NULL,
  bio_ref_interval json NOT NULL,
  comments TEXT NOT NULL,

  parameter_group_id UUID NOT NULL,
  parameter_rule_id UUID NOT NULL,
  FOREIGN KEY(parameter_group_id) REFERENCES parameter_groups(id),
  FOREIGN KEY(parameter_rule_id) REFERENCES parameter_rules(id)
);

CREATE TABLE IF NOT EXISTS reports (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

  parameter_group_id UUID NOT NULL,
  registration_id UUID NOT NULL,
  campaign_id UUID NOT NULL,
  FOREIGN KEY(registration_id) REFERENCES registrations(id),
  FOREIGN KEY(campaign_id) REFERENCES campaigns(id),
  FOREIGN KEY(parameter_group_id) REFERENCES parameter_groups(id)
);


