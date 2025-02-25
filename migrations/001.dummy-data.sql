INSERT INTO
    transactions (
        id,
        item_id,
        warehouse_id,
        kind,
        amount,
        description,
        cost,
        created_at
    )
VALUES
    (
        'b2f867bd-2101-4f74-aea7-99475b3066d1',
        'df4985ba-45ee-45dc-b31f-8fcbd677e9a2',
        '7d090868-3df5-44e7-9280-3cad6204be59',
        'input',
        10,
        'Ingreso 10 peras',
        10,
        date()
    ),
    (
        '99c78363-39d5-49b9-b21f-bf4db731542f',
        'df4985ba-45ee-45dc-b31f-8fcbd677e9a2',
        '7d090868-3df5-44e7-9280-3cad6204be59',
        'output',
        -2,
        'Salen 2 peras',
        10,
        date()
    );
