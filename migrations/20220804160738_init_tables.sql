-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    username varchar(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS public.authors (
    id SERIAL PRIMARY KEY,
    author VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS public.books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    author_id int,
    constraint fk_author
        FOREIGN KEY (author_id)
            REFERENCES public.authors(id)
                ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.reviews (
    id SERIAL PRIMARY KEY,
    rating int NOT NULL,
    review_text varchar(5000),
    "time" timestamp NOT NULL,
    book_id int,
    user_id int,
    constraint fk_book
        FOREIGN KEY (book_id)
            REFERENCES public.books(id)
                ON DELETE CASCADE,
    constraint fk_user
        FOREIGN KEY (user_id)
            REFERENCES public.users(id)
                ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.reviews;
DROP TABLE IF EXISTS public.books;
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.authors;
-- +goose StatementEnd
