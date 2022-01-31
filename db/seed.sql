DELETE FROM schedules;
DELETE FROM rooms;
DELETE FROM studios;
DELETE FROM schools;

INSERT INTO schools VALUE (NULL, "Studio Landin'", now(), now());
SELECT id INTO @school_id FROM schools WHERE name = "Studio Landin'";

INSERT INTO studios VALUE (NULL, "立川校", "東京都立川市", @school_id, now(), now()),
                          (NULL, "国分寺校", "東京都国分寺市", @school_id, now(), now());
SELECT id INTO @studio_id FROM studios WHERE name = "立川校";

INSERT INTO rooms VALUE (NULL, "A", 50, @studio_id, now(), now());
SELECT id INTO @room_id FROM rooms WHERE studio_id = @studio_id AND name = "A";

INSERT INTO schedules VALUE (NULL, 1, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 1, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 1, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 1, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 1, "21:45", "22:55",  @room_id, now(), now()),
                            (NULL, 2, "15:30", "16:40",  @room_id, now(), now()),
                            (NULL, 2, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 2, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 2, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 2, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 2, "21:45", "22:55",  @room_id, now(), now()),
                            (NULL, 3, "15:30", "16:40",  @room_id, now(), now()),
                            (NULL, 3, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 3, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 3, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 3, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 3, "21:45", "22:55",  @room_id, now(), now()),
                            (NULL, 4, "15:30", "16:40",  @room_id, now(), now()),
                            (NULL, 4, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 4, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 4, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 4, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 4, "21:45", "22:55",  @room_id, now(), now()),
                            (NULL, 5, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 5, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 5, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 5, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 5, "21:45", "22:55",  @room_id, now(), now()),
                            (NULL, 5, "13:00", "14:00",  @room_id, now(), now()),
                            (NULL, 6, "14:15", "15:25",  @room_id, now(), now()),
                            (NULL, 6, "15:30", "16:40",  @room_id, now(), now()),
                            (NULL, 6, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 6, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 6, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 6, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 6, "21:45", "22:55",  @room_id, now(), now()),
                            (NULL, 7, "14:15", "15:25",  @room_id, now(), now()),
                            (NULL, 7, "15:30", "16:40",  @room_id, now(), now()),
                            (NULL, 7, "16:45", "17:55",  @room_id, now(), now()),
                            (NULL, 7, "18:00", "19:10",  @room_id, now(), now()),
                            (NULL, 7, "19:15", "20:25",  @room_id, now(), now()),
                            (NULL, 7, "20:30", "21:40",  @room_id, now(), now()),
                            (NULL, 7, "21:45", "22:55",  @room_id, now(), now());