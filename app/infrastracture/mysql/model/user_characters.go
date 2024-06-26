// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserCharacter is an object representing the database table.
type UserCharacter struct {
	ID          uint64    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID      uint64    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CharacterID uint64    `boil:"character_id" json:"character_id" toml:"character_id" yaml:"character_id"`
	Level       int       `boil:"level" json:"level" toml:"level" yaml:"level"`
	Exp         int       `boil:"exp" json:"exp" toml:"exp" yaml:"exp"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *userCharacterR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userCharacterL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserCharacterColumns = struct {
	ID          string
	UserID      string
	CharacterID string
	Level       string
	Exp         string
	CreatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	CharacterID: "character_id",
	Level:       "level",
	Exp:         "exp",
	CreatedAt:   "created_at",
}

var UserCharacterTableColumns = struct {
	ID          string
	UserID      string
	CharacterID string
	Level       string
	Exp         string
	CreatedAt   string
}{
	ID:          "user_characters.id",
	UserID:      "user_characters.user_id",
	CharacterID: "user_characters.character_id",
	Level:       "user_characters.level",
	Exp:         "user_characters.exp",
	CreatedAt:   "user_characters.created_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var UserCharacterWhere = struct {
	ID          whereHelperuint64
	UserID      whereHelperuint64
	CharacterID whereHelperuint64
	Level       whereHelperint
	Exp         whereHelperint
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperuint64{field: "`user_characters`.`id`"},
	UserID:      whereHelperuint64{field: "`user_characters`.`user_id`"},
	CharacterID: whereHelperuint64{field: "`user_characters`.`character_id`"},
	Level:       whereHelperint{field: "`user_characters`.`level`"},
	Exp:         whereHelperint{field: "`user_characters`.`exp`"},
	CreatedAt:   whereHelpertime_Time{field: "`user_characters`.`created_at`"},
}

// UserCharacterRels is where relationship names are stored.
var UserCharacterRels = struct {
}{}

// userCharacterR is where relationships are stored.
type userCharacterR struct {
}

// NewStruct creates a new relationship struct
func (*userCharacterR) NewStruct() *userCharacterR {
	return &userCharacterR{}
}

// userCharacterL is where Load methods for each relationship are stored.
type userCharacterL struct{}

var (
	userCharacterAllColumns            = []string{"id", "user_id", "character_id", "level", "exp", "created_at"}
	userCharacterColumnsWithoutDefault = []string{"user_id", "character_id", "level", "exp"}
	userCharacterColumnsWithDefault    = []string{"id", "created_at"}
	userCharacterPrimaryKeyColumns     = []string{"id"}
	userCharacterGeneratedColumns      = []string{}
)

type (
	// UserCharacterSlice is an alias for a slice of pointers to UserCharacter.
	// This should almost always be used instead of []UserCharacter.
	UserCharacterSlice []*UserCharacter
	// UserCharacterHook is the signature for custom UserCharacter hook methods
	UserCharacterHook func(context.Context, boil.ContextExecutor, *UserCharacter) error

	userCharacterQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userCharacterType                 = reflect.TypeOf(&UserCharacter{})
	userCharacterMapping              = queries.MakeStructMapping(userCharacterType)
	userCharacterPrimaryKeyMapping, _ = queries.BindMapping(userCharacterType, userCharacterMapping, userCharacterPrimaryKeyColumns)
	userCharacterInsertCacheMut       sync.RWMutex
	userCharacterInsertCache          = make(map[string]insertCache)
	userCharacterUpdateCacheMut       sync.RWMutex
	userCharacterUpdateCache          = make(map[string]updateCache)
	userCharacterUpsertCacheMut       sync.RWMutex
	userCharacterUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userCharacterAfterSelectMu sync.Mutex
var userCharacterAfterSelectHooks []UserCharacterHook

var userCharacterBeforeInsertMu sync.Mutex
var userCharacterBeforeInsertHooks []UserCharacterHook
var userCharacterAfterInsertMu sync.Mutex
var userCharacterAfterInsertHooks []UserCharacterHook

var userCharacterBeforeUpdateMu sync.Mutex
var userCharacterBeforeUpdateHooks []UserCharacterHook
var userCharacterAfterUpdateMu sync.Mutex
var userCharacterAfterUpdateHooks []UserCharacterHook

var userCharacterBeforeDeleteMu sync.Mutex
var userCharacterBeforeDeleteHooks []UserCharacterHook
var userCharacterAfterDeleteMu sync.Mutex
var userCharacterAfterDeleteHooks []UserCharacterHook

var userCharacterBeforeUpsertMu sync.Mutex
var userCharacterBeforeUpsertHooks []UserCharacterHook
var userCharacterAfterUpsertMu sync.Mutex
var userCharacterAfterUpsertHooks []UserCharacterHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserCharacter) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserCharacter) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserCharacter) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserCharacter) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserCharacter) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserCharacter) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserCharacter) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserCharacter) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserCharacter) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userCharacterAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserCharacterHook registers your hook function for all future operations.
func AddUserCharacterHook(hookPoint boil.HookPoint, userCharacterHook UserCharacterHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userCharacterAfterSelectMu.Lock()
		userCharacterAfterSelectHooks = append(userCharacterAfterSelectHooks, userCharacterHook)
		userCharacterAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userCharacterBeforeInsertMu.Lock()
		userCharacterBeforeInsertHooks = append(userCharacterBeforeInsertHooks, userCharacterHook)
		userCharacterBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userCharacterAfterInsertMu.Lock()
		userCharacterAfterInsertHooks = append(userCharacterAfterInsertHooks, userCharacterHook)
		userCharacterAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userCharacterBeforeUpdateMu.Lock()
		userCharacterBeforeUpdateHooks = append(userCharacterBeforeUpdateHooks, userCharacterHook)
		userCharacterBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userCharacterAfterUpdateMu.Lock()
		userCharacterAfterUpdateHooks = append(userCharacterAfterUpdateHooks, userCharacterHook)
		userCharacterAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userCharacterBeforeDeleteMu.Lock()
		userCharacterBeforeDeleteHooks = append(userCharacterBeforeDeleteHooks, userCharacterHook)
		userCharacterBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userCharacterAfterDeleteMu.Lock()
		userCharacterAfterDeleteHooks = append(userCharacterAfterDeleteHooks, userCharacterHook)
		userCharacterAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userCharacterBeforeUpsertMu.Lock()
		userCharacterBeforeUpsertHooks = append(userCharacterBeforeUpsertHooks, userCharacterHook)
		userCharacterBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userCharacterAfterUpsertMu.Lock()
		userCharacterAfterUpsertHooks = append(userCharacterAfterUpsertHooks, userCharacterHook)
		userCharacterAfterUpsertMu.Unlock()
	}
}

// One returns a single userCharacter record from the query.
func (q userCharacterQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserCharacter, error) {
	o := &UserCharacter{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for user_characters")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserCharacter records from the query.
func (q userCharacterQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserCharacterSlice, error) {
	var o []*UserCharacter

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to UserCharacter slice")
	}

	if len(userCharacterAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserCharacter records in the query.
func (q userCharacterQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count user_characters rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userCharacterQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if user_characters exists")
	}

	return count > 0, nil
}

// UserCharacters retrieves all the records using an executor.
func UserCharacters(mods ...qm.QueryMod) userCharacterQuery {
	mods = append(mods, qm.From("`user_characters`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`user_characters`.*"})
	}

	return userCharacterQuery{q}
}

// FindUserCharacter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserCharacter(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*UserCharacter, error) {
	userCharacterObj := &UserCharacter{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_characters` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userCharacterObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from user_characters")
	}

	if err = userCharacterObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userCharacterObj, err
	}

	return userCharacterObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserCharacter) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no user_characters provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userCharacterColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userCharacterInsertCacheMut.RLock()
	cache, cached := userCharacterInsertCache[key]
	userCharacterInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userCharacterAllColumns,
			userCharacterColumnsWithDefault,
			userCharacterColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userCharacterType, userCharacterMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userCharacterType, userCharacterMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_characters` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_characters` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_characters` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userCharacterPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into user_characters")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userCharacterMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for user_characters")
	}

CacheNoHooks:
	if !cached {
		userCharacterInsertCacheMut.Lock()
		userCharacterInsertCache[key] = cache
		userCharacterInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserCharacter.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserCharacter) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userCharacterUpdateCacheMut.RLock()
	cache, cached := userCharacterUpdateCache[key]
	userCharacterUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userCharacterAllColumns,
			userCharacterPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update user_characters, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_characters` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userCharacterPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userCharacterType, userCharacterMapping, append(wl, userCharacterPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update user_characters row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for user_characters")
	}

	if !cached {
		userCharacterUpdateCacheMut.Lock()
		userCharacterUpdateCache[key] = cache
		userCharacterUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userCharacterQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for user_characters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for user_characters")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserCharacterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userCharacterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_characters` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userCharacterPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in userCharacter slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all userCharacter")
	}
	return rowsAff, nil
}

var mySQLUserCharacterUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserCharacter) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no user_characters provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userCharacterColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserCharacterUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userCharacterUpsertCacheMut.RLock()
	cache, cached := userCharacterUpsertCache[key]
	userCharacterUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userCharacterAllColumns,
			userCharacterColumnsWithDefault,
			userCharacterColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userCharacterAllColumns,
			userCharacterPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model: unable to upsert user_characters, could not build update column list")
		}

		ret := strmangle.SetComplement(userCharacterAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`user_characters`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_characters` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userCharacterType, userCharacterMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userCharacterType, userCharacterMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for user_characters")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userCharacterMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userCharacterType, userCharacterMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for user_characters")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for user_characters")
	}

CacheNoHooks:
	if !cached {
		userCharacterUpsertCacheMut.Lock()
		userCharacterUpsertCache[key] = cache
		userCharacterUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserCharacter record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserCharacter) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no UserCharacter provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userCharacterPrimaryKeyMapping)
	sql := "DELETE FROM `user_characters` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from user_characters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for user_characters")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userCharacterQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no userCharacterQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from user_characters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for user_characters")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserCharacterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userCharacterBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userCharacterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_characters` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userCharacterPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from userCharacter slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for user_characters")
	}

	if len(userCharacterAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserCharacter) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserCharacter(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserCharacterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserCharacterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userCharacterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_characters`.* FROM `user_characters` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userCharacterPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in UserCharacterSlice")
	}

	*o = slice

	return nil
}

// UserCharacterExists checks if the UserCharacter row exists.
func UserCharacterExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_characters` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if user_characters exists")
	}

	return exists, nil
}

// Exists checks if the UserCharacter row exists.
func (o *UserCharacter) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserCharacterExists(ctx, exec, o.ID)
}
