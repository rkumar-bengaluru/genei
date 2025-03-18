insert into roles(name) values 
    ('Doctor'),
    ('DIC Mysore'),
    ('Lab Tect'),
    ('Camp User'),
    ('System Admin'),
    ('doctor'),
    ('audio'),
    ('opto'),
    ('spyro');

insert into departments(name) values 
    ('IT Admin'),
    ('Call Coordinator'),
    ('Operations Manager'),
    ('Administrator'),
    ('Marketing & Sales'),
    ('HR'),
    ('Admin');

insert into states(name) values 
    ('Karnataka'),
    ('Andhra Pradesh'),
    ('Haryana'),
    ('Manipur'),
    ('Sikkim'),
    ('Tamil Nadu'),
    ('Meghalaya'),
    ('Himachal Pradesh'),
    ('Arunachal Pradesh'),
    ('Assam'),
    ('Jharkhand'),
    ('Mizoram'),
    ('Telangana'),
    ('Bihar'),
    ('Nagaland'),
    ('Tripura'),
    ('Chhattisgarh'),
    ('Kerala'),
    ('Odisha'),
    ('Uttarakhand'),
    ('Goa'),
    ('Madhya Pradesh'),
    ('Punjab'),
    ('Uttar Pradesh'),
    ('Gujarat'),
    ('Maharashtra'),
    ('Rajasthan'),
    ('West Bengal'),
    ('Chandigarh'),
    ('The Government of NCT of Delhi'),
    ('Andaman and Nicobar Islands'),
    ('Dadra and Nagar Haveli and Daman & Diu'),
    ('Jammu & Kashmir'),
    ('Lakshadweep'),
    ('Puducherry');

insert into districts(name, state_id) values 
    ('Bagalkote','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Bengaluru Urban','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Bengaluru Rural','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Belagavi','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Ballari','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Bidar','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Vijayapura','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Chamarajanagar','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Chikkaballapura','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Chikkamagaluru','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Chitradurga','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Dakshina Kannada','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Davanagere','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Dharwad','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Gadag','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Kalaburagi','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Hassan','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Haveri','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Kodagu','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Kolar','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Koppal','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Mandya','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Mysuru','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Raichur','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Ramanagara','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Shivamogga','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Tumakuru','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Udupi','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Uttara Kannada','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Vijayanagara','35e89f12-dba7-406e-8be8-c515e02b71b5'),
    ('Yadgiri','35e89f12-dba7-406e-8be8-c515e02b71b5');

insert into pin_codes(pin_code, post_office, district_id) values 
    ('587111','Achanur B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587201','Adagal B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587330','Adihudi B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587206','Agasarkoppa B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587124','Aihole B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587301','Algur B.O' ,'0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587155','Alur S.K. B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587121','Amalzeri B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587118','Amaravati B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587125','Amarawadgi B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587112','Amingad S.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587116','Anagwadi B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587206','Anawal B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587116','Arakeri B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587311','Asangi B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587201','Bachingudd B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587116','Badagandi B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587117','Badagi B.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587201','Badami Extension S.O','0303f494-ad0a-4918-b943-24e97aceef46'),
    ('587201','Badami S.O','0303f494-ad0a-4918-b943-24e97aceef46');

-- for kolar
insert into pin_codes(pin_code, post_office, district_id) values 
    ('563161','Addagal B.O','75f1be57-bdf3-49bb-a671-3dd05a08e665'),
    ('563138','Addagal B.O','75f1be57-bdf3-49bb-a671-3dd05a08e665');
-- work orders
insert into work_orders(name) values 
    ('VaraLakshmi | Kolar 42/2024-25| Kolar');

-- application access
insert into application_access(name) values 
    ('ALL');
-- assigning_authority
insert into assigning_authority(name) values 
    ('ALL');
-- stores
insert into stores(name) values 
    ('ALL');
-- campaigns
insert into campaigns(district_id, state_id, estimated_target_screening,pin_code_id,
                      taluk_name, application_access_id, camp_name,screening_start_date,
                      screening_start_time,assigning_authority_id,store_id,work_order_id,
                      latitude,longitude,labour_inspector_name,union_name,union_leader_name,description) values 
    ('75f1be57-bdf3-49bb-a671-3dd05a08e665','35e89f12-dba7-406e-8be8-c515e02b71b5',75,'f65e6e2b-1491-4d78-8611-8fe3e6c38604',
    'Srinivaspur','da3bd164-0e06-46d8-89c6-7f22e370ebad','Mastenahalli','03-18-2025',
    '15:00','e37a717f-8a5a-432c-9796-2ccc29d7b357','aaafc79f-993c-4400-9ba8-729afa8123ef','2cac6a45-1e8a-4d82-84ed-8ba196fef073',
    '13.278457904360922','78.11936391384651','','','','')

