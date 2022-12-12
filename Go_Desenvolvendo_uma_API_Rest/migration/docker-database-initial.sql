create table books(
    id serial primary key,
    title varchar,
    author varchar,
    pages integer
);

INSERT INTO books(title, author, pages) VALUES
('Se quiser mudar o mundo: Um guia político para quem se importa', 'Sabrina Fernandes', 192);
('Terra das mulheres', 'Charlotte Perkins Gilman', 256);