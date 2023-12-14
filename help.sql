SELECT COUNT(*) FROM image;
SELECT SUM(id) FROM image;
show tables;
select * from task;
select * from avif join task on task.id = avif.task_id;
select * from task;
select * from avif;
drop table av,custom,err,file,image,save,telegraph,text,video,audio,avif,task;
show tables;