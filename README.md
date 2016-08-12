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
            "group_id": "int64",
            "pk": "user_id"
        },
        "queries": [
            "findById": "SELECT * FROM users WHERE user_id = ?",
            "findManyByGroupId: "SELECT * FROM users WHERE group_id = ?"
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
}

func (u *User) IsStored() bool {
    return u.UserId > 0
}

func (u *User) Scan(s gen.ScanFunc) error {
    return s(
        &u.UserId,
        &u.Name,
        &u.CreateDate,
        &u.UpdateDate,
        &u.Telephone,
        &u.GroupId,
    )
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

func FindManyByGroupId(groupId int64) ([]*User, error) {

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
