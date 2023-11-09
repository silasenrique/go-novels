-- name: CreateSite :exec
insert into site (id, name, url, creation_date) values (?, ?, ?, ?);

-- name: ListSite :many
select * from site s where (id = ? or ? = 0) order by id;

-- name: DeleteSite :exec
delete from site where id = ?;

-- name: CreateNovel :exec
insert into novel (id, name, creation_date) values (?, ?, ?);

-- name: DeleteNovel :exec
delete from novel where id = ?;

-- name: ListNovels :many
select * from novel n where (id = ? or ? = 0) order by id;

-- name: ListNovelSites :many
select * from novel_site l where (l.id_novel = ? or ? = 0) order by id;

-- name: ListSiteNovels :many
select * from novel_site l where (l.id_site = ? or ? = 0) order by id;
