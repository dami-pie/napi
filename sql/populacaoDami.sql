USE dami;

INSERT INTO zones (zoneId, zoneDesc)
VALUES (1, "Bloco K"),
(2, "Bloco B"),
(3, "Bloco I"),
(4, "Bloco J"),
(5, "Bloco C"),
(6, "Salas Professores");



INSERT INTO room (roomNumber, zoneId, roomBuild)
VALUES (1, 6, "P"),
(2, 6, "P"),
(3, 6, "P"),
(1, 1, "K"),
(2, 1, "K"),
(3, 1, "K"),
(1, 2, "B"),
(2, 2, "B"),
(3, 2, "B"),
(1, 3, "I"),
(2, 3, "I"),
(3, 3, "I"),
(1, 4, "J"),
(2, 4, "J"),
(3, 4, "J"),
(1, 5, "C"),
(2, 5, "C"),
(3, 5, "C");


INSERT INTO userGroup (userGroupId, userGroupDesc)
VALUES (1, "Professores"),
(2, "Funcionarios"),
(3, "AlunosGraduacao"),
(4, "AlunosMestrado");


INSERT INTO accessPermit (accessPermitDesc, userGroupId, zoneId)
VALUES ("Acesso prof sala prof", 1, 1),
("Acesso prof sala aula", 1, 2),
("Acesso prof sala aula", 1, 3),
("Acesso prof sala aula", 1, 4),
("Acesso prof sala aula", 1, 5),
("Acesso prof sala aula", 1, 6),
("Acesso Limpeza", 2, 1),
("Acesso Limpeza", 2, 2),
("Acesso Limpeza", 2, 3),
("Acesso Limpeza", 2, 4),
("Acesso Limpeza", 2, 5),
("Acesso Limpeza", 2, 6),
("Acesso Alunos sala aula", 3, 1),
("Acesso Alunos sala aula", 3, 2),
("Acesso Alunos sala aula", 3, 3),
("Acesso Alunos sala aula", 3, 4),
("Acesso Alunos sala aula", 3, 5),
("Acesso Alunos sala mestrado", 4, 1);


INSERT INTO userdata (userEmail, userGroupId)
VALUES ("bruno@ecomp.com.br", 1),
("alexandre@ecomp.com.br", 1),
("luis@ecomp.com.br", 1),
("tarciana@ecomp.com.br", 1),
("rativa@ecomp.com.br", 1),
("tio@ecomp.com.br", 2),
("limpeza@ecomp.com.br", 2),
("gyver@ecomp.com.br", 2),
("azevedo@ecomp.com.br", 3),
("marta@ecomp.com.br", 3),
("murilo@ecomp.com.br", 3),
("helio@ecomp.com.br", 3),
("will@ecomp.com.br", 3),
("escorel@ecomp.com.br", 3),
("erick@ecomp.com.br", 3),
("glauco@ecomp.com.br", 3),
("thales@ecomp.com.br", 3),
("heleno@ecomp.com.br", 4),
("furilo@ecomp.com.br", 4),
("popers@ecomp.com.br", 4),
("uiu@ecomp.com.br", 4);


INSERT INTO userCard (cardEnable, userID, cardHash)
VALUES (1, 2, "2808"),
(0, 1, "1351"),
(1, 3, "4790"),
(1, 4, "1083"),
(1, 5, "3973");



INSERT INTO entranceLog (roomNumber, roomBuild, entranceType, userID, dateHour, cardHash)
VALUES (2,"P", 0, 1, '2023-06-14',NULL),
(3,"I", 1, 2, '2023-06-14',"4790"),
(1,"J", 1, 2, '2023-06-14',"1083"),
(2,"C", 1, 3, '2023-06-14',"3973"),
(3,"B", 1, 5, '2023-06-14',"2808"),
(2,"I", 0, 1, '2023-06-14',NULL),
(3,"K", 0, 2, '2023-06-14',NULL),
(1,"C", 0, 7, '2023-06-14',NULL),
(3,"B", 0, 8, '2023-06-14',NULL),
(2,"K", 0, 4, '2023-06-14',NULL),
(3,"J", 0, 6, '2023-06-14',NULL),
(3,"J", 0, 3, '2023-06-14',NULL),
(2,"I", 0, 1, '2023-06-14',NULL),
(2,"B", 0, 9, '2023-06-14',NULL),
(1,"K", 0, 1, '2023-06-14',NULL),
(1,"I", 0, 2, '2023-06-14',NULL),
(2,"K", 0, 3, '2023-06-14',NULL),
(3,"C", 0, 4, '2023-06-14',NULL),
(3,"B", 0, 4, '2023-06-14',NULL),
(1,"K", 0, 6, '2023-06-14',NULL),
(2,"J", 0, 2, '2023-06-14',NULL),
(2,"J", 0, 1, '2023-06-14',NULL),
(1,"I", 0, 3, '2023-06-14',NULL),
(2,"B", 0, 1, '2023-06-14',NULL),
(3,"K", 0, 5, '2023-06-14',NULL);
