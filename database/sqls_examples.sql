--installers
INSERT INTO pcam.installers (fullname, alias, username, usertype, status, "password", created_at, updated_at) VALUES('', '', '', '', 0, '', '', '');

--clients
INSERT INTO pcam.clients (fullname, alias, status, created_at, updated_at, codclient) VALUES('', '', 0, '', '', '');

--cams
INSERT INTO pcam.cams (devicename, alias, streamtype, urlcamstream, status, created_at, updated_at) VALUES('', '', '', '', 0, '', '');

--users
INSERT INTO pcam.users (fullname, alias, username, usertype, status, "password", created_at, updated_at) VALUES('', '', '', '', 0, '', '', '');

--licences
INSERT INTO pcam.licences (uniccode, clientid, installerid, licencetype, created_at, updated_at) VALUES('', 0, 0, '', '', '');
