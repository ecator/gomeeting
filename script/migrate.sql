--- migrate 0.x -> 1.x

use gomeeting;

ALTER TABLE meeting ADD id VARCHAR(32) FIRST;

ALTER TABLE meeting CHANGE make_date create_time INT UNSIGNED NOT NULL ;

UPDATE meeting SET id = MD5( CONCAT( room_id, maker, start_time, end_time, create_time ) );

ALTER TABLE meeting ADD PRIMARY KEY ( id) ;


UPDATE meeting 
    SET start_time = UNIX_TIMESTAMP( STR_TO_DATE( create_time, '%Y%m%d' ) ) + start_time *60 -8 *60 *60
       ,end_time = UNIX_TIMESTAMP( STR_TO_DATE( create_time, '%Y%m%d' ) ) + end_time *60 -8 *60 *60
       ,create_time = UNIX_TIMESTAMP( STR_TO_DATE( create_time, '%Y%m%d' ) );