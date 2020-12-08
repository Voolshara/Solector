# Импорт библиотеки
import sqlite3

data = [['HVL-395/HJT', 18490, 395, 19.75, 'https://www.hevelsolar.com/loaded/catalog/goods/0aba2318141e63ca08b2ce7fac84d97d.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-395hjt/'],
        ['HVL-390/HJT', 18290, 390, 19.5, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/1e0f6c8ea7310cb04c09bd64585b1dbf.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-390hjt/'],
        ['HVL-385/HJT', 17990, 385, 19.3, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/5b53904ed9cd5f236149a8ac690a391b.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-385hjt/'],
        ['HVL-380/HJT', 17790, 380, 19, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/a67b37db6ad01447c822b8456735ea23.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-380hjt/'],
        ['HVL-375/HJT', 17590, 375, 18.75, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/604097ab1566a473899ffc99ecca3bea.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-375hjt/'],
        ['HVL-370/HJT', 17390, 370, 18.5, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/9b6aad14499577f19ac2d0fc1e193618.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-370hjt/'],
        ['HVL-365/HJT', 17090, 365, 18.25, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/215b0bed55c475d34426d9dc424ded6c.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-365hjt/'],
        ['HVL-360/HJT', 16890, 360, 18, 'https://www.hevelsolar.com/loaded/catalog/goods/c0a68703e15a31c6023882eb93da4c24.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-360hjt/'],
        ['HVL-330/HJT', 15190, 330, 19.7, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/952a2d8b9b474c4b944017296059fa8c.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-330hjt/'],
        ['HVL-325/HJT', 14890, 325, 19.4, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/6c68e7070956697c50acbef7ce409fc3.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-325hjt/'],
        ['HVL-320/HJT', 14690, 320, 19.1, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/c0d67c470705d50255e7fe2a18aebdf5.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-320hjt/'],
        ['HVL-315/HJT', 14490, 315, 19.04, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/5aaeb9180bf2bf03883f1ddd2fb8a8e2.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-315hjt/'],
        ['HVL-310/HJT', 14290, 310, 18.76, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/d8abfc5093c85171a3f493eaaab4ce62.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-310hjt/'],
        ['HVL-290/HJT', 11790, 290, 17.32, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/2b0e51dd36d1f49671cc6c5755f5bed8.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-290hjt/'],
        ['HVL-250/HJT', 10190, 250, 14.93, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/bda26942eda46c5e2020d255ec1f64c8.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-250hjt/'],
        ['HVL-125/O', 4625, 125, 8.74, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/39a6d1c0e23a39d94e1634d81fbcb95e.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-125o/'],
        ['HVL-120/O', 4440, 120, 8.39, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/2b50dda52fb68e3f53073a9c3244a69d.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-120o/'],
        ['HVL-115/O', 4255, 115, 7.9, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/217d0489dbc4b3030d54fbc0f228988f.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-115o/'], #Error pers
        ['HVL-105/O', 3885, 105, 7.34, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/6e1e3514129ed0baba18d3cd89606927.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-105o/'],
        ['HVL-300/HJT', 12190, 300, 17.98, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/32524347c166bbd0dc9d3d0752014186.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-300hjt/'],
        ['HVL-270/HJT', 10990, 270, 16.12, 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/085770ce2086d0c1b5c7675d52b181a9.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-270hjt/'],
        ['HVL-260/HJT', 10590, 260, 15.53 , 'https://www.hevelsolar.com/loaded/catalog/goods/thumbs/84def46e327046071a6e8ff3d7357d45.png', 'https://www.hevelsolar.com/catalog/solnechnye-moduli/modul-fotoelektricheskii-hvl-260hjt/']]

# Подключение к БД
con = sqlite3.connect("panels.sqlite")

# Создание курсора
cur = con.cursor()
i = 1
for elem in data:
        id = i
        name = elem[0]
        cost = elem[1]
        power = elem[2]
        efficiency = elem[3]
        imgLink = elem[4]
        companyName = "Hevel Solar"
        panelLink = elem[5]

        # Выполнение запроса и получение всех результатов
        # cur.execute("""INSERT INTO genres(id,name, cost, power, efficiency, imgLink, companyName, panelLink)
        #  VALUES(""" + id + """, """ + name + """, """ + cost + """, """ + power + """, """ + efficiency + """, """ + imgLink + """, """ + companyName + """, """ + panelLink + """)""")
        # i += 1

        cur.execute("""INSERT INTO panels(id,name, cost, power, efficiency, imgLink, companyName, panelLink)
                 VALUES(?, ?, ?, ?, ?, ?, ?, ?)""", (id, name, cost, power, efficiency, imgLink, companyName, panelLink,))
        con.commit()
        i += 1
        print(elem)

con.close()
