# Payment_System

```
git clone ...

cd .\Payment_System\

go run .\cmd\main.go

go test .\pkg\PaymentSystem\
```

## Пример вывода:

```
Account BY13NBRB3600900000002Z00AB00 (EMITION BYN) created
Account BY13NBRB3600900000002Z00AB01 (EMITION USD) created
Account BY13NBRB3600900000002Z00AB10 (WITHDROWAL BYN) created
Account BY13NBRB3600900000002Z00AB11 (WITHDROWAL USD) created
Эмиссионный счет в BYN:  BY13NBRB3600900000002Z00AB00
Счет для вывод денежных средств BYN из оборота BY13NBRB3600900000002Z00AB10
Emition for 1000.00 BYN complited
Emition for 1000.00 USD complited
Account BY13NBRB3600900000002Z00AB22 (MAZ BYN) created
Transfer for 100.00 BYN from BY13NBRB3600900000002Z00AB00 (EMITION) to BY13NBRB3600900000002Z00AB22 (MAZ) complited
Account BY13NBRB3600900000002Z00AB33 (EPAM USD) created
Account BY13NBRB3600900000002Z00AB44 (EPAM BYN) created
Transfer for 10.00 BYN from BY13NBRB3600900000002Z00AB22 (MAZ) to BY13NBRB3600900000002Z00AB44 (EPAM) complited
Transfer for 20.00 BYN from BY13NBRB3600900000002Z00AB22 (MAZ) to BY13NBRB3600900000002Z00AB44 (EPAM) complited
Start Withdrowal ... Transfer for 30.00 BYN from BY13NBRB3600900000002Z00AB22 (MAZ) to BY13NBRB3600900000002Z00AB10 (WITHDROWAL) complited
Текущий баланс  MAZ 40

Список счетов:
[
  {
    "iban": "BY13NBRB3600900000002Z00AB00",
    "currency": "BYN",
    "owner": {
      "uuid": "ed5d072b-5879-4d42-a41f-0f9e7aec6174",
      "name": "EMITION"
    },
    "balance": 900,
    "status": "active"
  },
  {
    "iban": "BY13NBRB3600900000002Z00AB01",
    "currency": "USD",
    "owner": {
      "uuid": "e5b38048-ceba-4ada-853f-7cf25940c26e",
      "name": "EMITION"
    },
    "balance": 1000,
    "status": "active"
  },
  {
    "iban": "BY13NBRB3600900000002Z00AB10",
    "currency": "BYN",
    "owner": {
      "uuid": "6452be2b-ed9b-44a1-a02e-b752cfae5ee2",
      "name": "WITHDROWAL"
    },
    "balance": 30,
    "status": "active"
  },
  {
    "iban": "BY13NBRB3600900000002Z00AB11",
    "currency": "USD",
    "owner": {
      "uuid": "d7fe5ae4-6f58-4861-9257-6c9b67d454ef",
      "name": "WITHDROWAL"
    },
    "balance": 0,
    "status": "active"
  },
  {
    "iban": "BY13NBRB3600900000002Z00AB22",
    "currency": "BYN",
    "owner": {
      "uuid": "6261d2b1-6b8e-464c-b771-9be79c5cde51",
      "name": "MAZ"
    },
    "balance": 40,
    "status": "active"
  },
  {
    "iban": "BY13NBRB3600900000002Z00AB33",
    "currency": "USD",
    "owner": {
      "uuid": "584ea98e-a291-441b-9f59-891487e70334",
      "name": "EPAM"
    },
    "balance": 0,
    "status": "active"
  },
  {
    "iban": "BY13NBRB3600900000002Z00AB44",
    "currency": "BYN",
    "owner": {
      "uuid": "7dcb1dd3-16d1-4faf-91e5-9eee9a182d56",
      "name": "EPAM"
    },
    "balance": 30,
    "status": "active"
  }
]

```
