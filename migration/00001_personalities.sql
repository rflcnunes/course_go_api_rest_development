create table personalities(
    id serial primary key,
    name varchar,
    history varchar
);
INSERT INTO personalities(name, history)
VALUES (
        'Albert Einstein',
        'Albert Einstein was a German-born theoretical physicist who developed the theory of relativity, one of the two pillars of modern physics.'
    ),
    (
        'Marie Curie',
        'Marie Curie was a Polish-born physicist and chemist who conducted pioneering research on radioactivity. She was the first woman to win a Nobel Prize and the only person to win Nobel Prizes in two different scientific fields.'
    ),
    (
        'Nelson Mandela',
        'Nelson Mandela was a South African anti-apartheid revolutionary and politician who served as President of South Africa from 1994 to 1999. He was the country''s first black head of state and the first elected in a fully representative democratic election.'
    ),
    (
        'Ada Lovelace',
        'Ada Lovelace was an English mathematician and writer, chiefly known for her work on Charles Babbage''s early mechanical general-purpose computer, the Analytical Engine. She is often regarded as the first computer programmer.'
    );