-- Taking the guard for a walk
with recursive guard_steps as (
    select y, x, -1 as dir_y, 0 as dir_x
    from day06
    where char = '^'
    union all
    select
        guard_steps.y + dir_y,
        guard_steps.x + dir_x,
        -- turn right: swapping y and x, switching sign on y
        case when facing.char = '#' then dir_x else dir_y end as dir_y,
        case when facing.char = '#' then -dir_y else dir_x end as dir_x
    from guard_steps
        left outer join day06 facing
            on facing.y = guard_steps.y + (dir_y * 2)
            and facing.x = guard_steps.x + (dir_x * 2)
    where facing is not null
    )

-- Part 1
select count(distinct y * 1000 + x) + 1 -- +1 since we're exiting recursion before just before the last step
from guard_steps;

/* setup
drop table if exists day06;
drop table if exists day06_example;

create table day06 (y int, x int, char char(1), primary key(char, y, x));
create table day06_example (y int, x int, char char(1), primary key(char, y, x));

insert into day06
select
    y,
    row_number() over(partition by y order by 1) as x,
    char
from (
    select
       row_number() over(order by 1) as y,
       regexp_split_to_array(row, '') as chars
    from unnest(string_to_array(
'..........#...........#...........#.............................#......................#.........#................................
..........................#.........#......#.................................#...........................#........................
..................................#...#................................................................................#..........
.#...............#.....#..........#.........#......#........#................................#......#.............................
...#.......#.............#.............................................#.........##...............#.................#.............
.......................................#.............#..................#........#.......#.........#...............#.........#....
.....#..................#..#..#.........................#...................#....#........#.............................#.........
.......#.............................#..........#......#..........................##......#.......................##..............
.....#.....#........................................................................................#.......#.....................
.................................#............#...#.....................................................#.........#...............
...............#.....#..#....#....#.....................................................#............#........#.#.................
.......#.....#...........................................................#............................#.................#....#....
...........................#.................#....................................#...............#..#........................#...
.....#.........................................#.................#...##....................................#......................
.............#...#..................#......#.................#..........................................#.....#..........#........
.................#...............#............#.......#.#...............................................#..................#......
.....#..#.....................#.......................................................#..............#..............#.......##....
..##........#....................................................#.............#..............#........................#..#.......
..............................#...........#...#........................#............#..........................#........#........#
.....................................##...............................#............#............#....#............................
...##........................#.....#....#.........#.................................................#.............................
...............................#................................#...............#.....#.....................#.....................
................................................#..........#....#............................#.....#.......#..................#...
.............#.....#.........#.#....#................##......#.............#......................................................
..........#..............#..........#......#.........#...#.........................................#.#.....#.........#............
..................#......................#.............................#...................#.....................................#
................#.....#..............................#.............#.............#..#.....#....................................#..
.........#.#................#.........#.................#...................#................................#................#...
.............................................................#..............#.....................................................
......#.........................#...........#.........#..................#....................................#..#...........##...
.............................................................................#.............#........................#...#.........
......#......#..#............#..............................................#....#..#..#................#.........#....#..........
........#........#.........#....................................................^..#...#....................#...............#.....
..#....#...#.......................................................................................................#..............
.....#...........................................#..............#.........#....#........................................#.........
.....................................#..............#............................................................#................
.....#.......................#........................#..........#.............#.................#..................#.............
.......................................#......#.....................................#.........................#...................
...........................................................................................#............................##........
..#............#..........................#...........#...#..................#......#.................#.#.................#......#
.###......#....................#...#...........................................................#.........#.....#..............#.#.
....................................#..........#.....#.......................#...............#.....#......#...................#...
...#......#..................#......#.......#..............#..............#.......................................................
..........#....................................................................#................##.#...#..........................
.........................................................#..................##.....................#..##..........................
..................#......#.................................#.................#.##................................................#
...#..........#...#..#..............................................#...........................##........#.......................
.....#................................................#...........#..................#...#.................................#......
.......#.....#.........................#..........................................................................................
....................................................#........#....#....................................##..............#.....#....
............#........#.........................#.................................#..........................................#...#.
.....................................#........................#...................#......#.#...........#..#..........#...........#
................#...............#.........................#.........................#........#..#.................................
.............#.........#.............#.#.#.........#...................#.......#.......................#..........................
.............#......................................................#................................................#.......#....
....#....................................................................#.#........#.................#..........................#
.................................#......#....................................................................#.....#..............
............##...................#.......................................................................#.........#........#.....
#..................................................#..#...............#............#.............#...............................#
.............................#.................................#.....#..........................#.............#..#...#...#........
...#.....#..............................#.#......................................#................................................
...#............##...............................#............#..............................#................................#...
#.................#......##..............................................................#...........#........#...#...............
....##.....................................................................................................#.#....................
#.................#..................................................#..................................................#.........
.........#.........#...........................#.................#.#.............................#................................
................................#............#................................................................#...................
.....................................#............................................#.##......##........#...........................
....#....................................................#...#.#.#....................................................#...........
.........................#.#...................#...................##...............#........................#....................
..................................#...................................................................#.......#..................#
...................#.............................................#.......................#.#.....................#....#...........
#....................................#.#..#.#..........#.........#.........##.........#......................#............#.......
...............................................................#.............................##....#.......#......................
.............#......#.#.........#............................#............................#.......................................
......................#.........#.........................................................#...#............#......................
...................#.....#.....................................................#.....................................#............
...............................#.....#..................#.......##................................................................
.............................................#.........................#...........#..............................#..#............
#..##..................................................................................................#......#...........#.....#.
.......................................................#.....................#.............................#......................
...........................................................................................................#......................
......................#..............................................................#............................................
...........................#.....#......................................#.........#........#..#.#.................................
.#.......#........................................................................................................................
.#.......#...........................................#................................#...........................................
.........................................................#........................................#......#........................
...............................#...#........#............................................#.................#......................
......#.........................#......................#.....#.#..............#..................#.......#..................#.....
.................#........#...#........................#.............................#............................................
#...........................................................................#............................#................#.......
..........................................................................#................#......................................
.............#.............................................................................#.........................#............
.#...........#.......#.......#.....#....#.......#........#....................#........................#........................#.
...#..............#.............................#.............................................................#...#..#..........#.
......#......#.#...............##................................................................................................#
....#........#.##......#...#......................................................................................................
..#..#.#.#.....................#................................#...................................#.....#...#...................
.....................#.#........#..............##..............#.................#..............................................#.
...................................................................................#........#.......#....#............#...........
.....#....#..................#...............#.............................................#.#....................#...............
.....................#..........#............#..................................#......................#.#.....#................#.
..........#.........................................#...............#.............................................................
.......#......................#.............#.....#..............#........................#.....#.........#.................#...#.
..................................................................#.....#...............................................#........#
.......#..........#.......#....#.......#..........#....#............................................#.#...........................
..#...........#..............................#....................#....................#...#......................................
.......#...............................#............#.....#.....................#.................................................
.........#..................#..................#..................................................................................
..........................................................................................#....#...................#.......##.....
..................#.............................##.......................................................#.......#......#......#..
...#.......#..............................................................................................#.......................
............#..#..............................................................#...........##....#....#.#................#........#
............#.........................#........#....#..#..................................#..................#....................
......##......................#...##..............................#...............................................................
........................#................#..............#...##...............................#..................##....#...........
....#.#...#.....#..................................#..............................#...........#.....................#.............
......................#.........................................................................#..........#.............#........
.#......#............................................................#...............................#...#........................
.......##.....................#.....#...........#......##.......##...................#...............................#............
...#...#........#...............#......#......#.....#...#.................................................................#.......
..................#.......................#...............................................#.................#.............#..#....
.....................#.......#....#...#............................#.....................#............#.....................#.....
................................................................#....#......#............#..................#.........#...........
.......................#...................#..............................................................................###.....
...............................#...........#................................#........#..#....#............................#.......
....#......#...............................#...............#......................................#...............................
...........#...........................##.........................................................................#.#.............
.........................#.........#.........................................#.....#...................#.....#.......#............
.............................#.....#................##..#.....................#...................................................'
        ,E'\r\n')) as row), unnest(chars) as char;

insert into day06_example
select
    y,
    row_number() over(partition by y order by 1) as x,
    char
from (
    select
       row_number() over(order by 1) as y,
       regexp_split_to_array(row, '') as chars
    from unnest(string_to_array(
'....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...'
        ,E'\r\n')) as row), unnest(chars) as char;
 */