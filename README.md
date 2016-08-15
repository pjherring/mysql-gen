This package will translate a set of json files into to a record struct and various findOne/findMany/count/store/delete methods

Given the file users.json

    {
        "name": "users",
        "fields": {
            "user_id": "int64",
            "name": "string",
            "create_date": "mysql.NullTime",
            "update_date": "mysql.NullTime",
            "telephone": "string",
            "group_id": "int64"
        },
        "queries": [
            "findById": "SELECT * FROM users WHERE user_id = ?",
            "findManyByGroupId: "SELECT * FROM users WHERE group_id = ? LIMIT ? OFFSET ?",
            "getManyUserIdsByName": "SELECT user_id FROM users WHERE name = ?"
        ]
    }

mysql-gen users.json model

Should produce the following in the users package inside of the model folder

in record_gen.go

    package users

    type User struct {
        UserId int64
        Name string
        CreateDate mysql.NullTime
        UpdateDate mysql.NullTime
        Telephone string
        GroupId int64
        IsStored bool
    }

    func (u *User) Scan(s gen.ScanFunc) error {
        err := s(
            &u.UserId,
            &u.Name,
            &u.CreateDate,
            &u.UpdateDate,
            &u.Telephone,
            &u.GroupId,
        )
        
        if err != nil {
            u.IsStored = true
        }

        return err
    }

in table_gen.go

    func (u *User) Store() error {
        if u.IsStored() {
            return u.insert()
        }

        return u.update()
    }

    func (u *User) update() error {
        _, err := gen.GetDb().Exec(`
            UPDATE users SET name = ?, create_date = ?, update_date = ?, telephone = ?, group_id = ?
            WHERE user_id = ?
        `)

        if err != nil {
            u.IsStored = true
        }

        return err
    }

    func (u *User) insert() error {
        r, err := gen.GetDb().Exec(`
            INSERT INTO users
            (name, create_date, update_date, telephone, group_id)
            VALUES
            (?, ?, ?, ?, ?)
        `)
        
        if err == nil {
            u.UserId = r.LastInsertId()
            u.IsStored = true
        }

        return err
    }

    func FindById(userId int64) (*User, error) {
        
        r := db.QueryRow(
            "SELECT * FROM users WHERE user_id = ?",
            userId,
        )

        retval := new(User)
        err := retval.Scan(r.Scan)

        return retval, err
    }

    func FindManyByGroupId(groupId int64, limit int, order int) ([]*User, error) {

        var retval []*User

        r, err := db.Query(
            "SELECT * FROM users WHERE group_id = ?",
            groupId,
        )

        if err != nil {
            return nil, err
        }

        defer r.Close()
        for r.Next() {
            u := new(*User)
            if err := u.Scan(r.Scan); err != nil {
                return nil, err
            }

            retval = append(retval, u)
        }

        return retval, nil
    }


    //"getManyUserIdsByName": "SELECT user_id FROM users WHERE name = ?"
    func GetManyUserIdsByName(name string) []int64 {
        var retval []int64

        r, err := db.Query(
            "SELECT user_id FROM users WHERE name = ?",
            name,
        )

        if err != nil {
            return nil, err
        }

        defer r.Close()
        for r.Next() {
            var i int64
            if err := r.Scan(&i); err != nil {
                return nil, err
            }

            retval = append(retval, i)
        }

        return retval
    }

