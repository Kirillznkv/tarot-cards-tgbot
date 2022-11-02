-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
insert into images (name, url, description) (
    values (
        'Шут',
        'https://i0.wp.com/all-tarot.ru/wp-content/uploads/2020/01/lightseers-00-fool-tarot-meaning.png?resize=406%2C700&ssl=1',
        'Светлое пророчество: Новые начинания, потенциал, приключения, энтузиазм, пробуждение, невинность, оптимизм.\nТеневое пророчество: Самонадеянность, ошибочное мнение, что ответ у вас уже есть, импульсивность и скоропалительность решений, недостаток опыта, глупость, аналитический паралич (чрезмерное анализирование ситуации).'
    ),
    (
        'Шут2',
        'https://i0.wp.com/all-tarot.ru/wp-content/uploads/2020/01/lightseers-01-magician-tarot-meaning.png?resize=406%2C700&ssl=1',
        'adfs'
    )
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
delete from images
    where name = 'Шут' or name = 'Шут2';
-- +goose StatementEnd
