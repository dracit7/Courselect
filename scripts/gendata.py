
# This scripts could generate some test data for the
# courselect database.

import itertools
import random
import time

def rand_time():
  stimes = ['08:00:00', '10:10:00', '14:00:00', '16:10:00', '18:30:00']
  etimes = ['09:40:00', '11:50:00', '15:40:00', '17:30:00', '21:50:00']
  index = random.randint(0, len(stimes)-1)
  return stimes[index], etimes[index]

def rand_id(prefix, grade, i):
  return "%s%d%d" % (prefix, grade, i+10000)

def rand_grade():
  return random.randint(2000, 2020)

def rand_name():
  surnames = ['张', '李', '林', '刘', '王', '黄', '宋', '丁', '周']
  names = ['豫','章','故','郡','洪','都','新','府','星','分','翼','轸','地','接','衡','庐','襟','三','江','而','带','五','湖','控','蛮','荆','而','引','瓯','越','物','华','天','宝','龙','光','射','牛','斗','之','墟','人','杰','地','灵','徐','孺','饯','子']
  return random.choice(surnames) + random.choice(names)

def rand_pw():
  pw = ''
  for _ in range(random.randint(8, 12)):
    pw += random.choice('123456789qwertyuiopasdfghjklzxcvbnm')
  return pw

# Amount of random data
major_num = 100
student_num = 10000
faculty_num = 1000
course_num = 100

# Output
output_file = 'scripts/testdata.sql'
f = open(output_file, 'w')

# Hint
hint = '''
/* This script generates some test data for the 
 * courselect database and should be run by the
 * courselect user like:
 *
 * mysql courselect -u courselect -p < testdata.sql
 * 
 * Attention: running this script may add
 * irrelevant data to your courselect database. */
'''
f.write(hint + '\n')

# Major information
majors = [
  '计算机科学与技术',
  '光学与电子信息',
  '船舶与海洋工程',
  '临床医学',
  '土木工程',
]
major_cnt = 1
for major in majors:
  f.write(
    "insert into major (name) values ('%s');\n" % major +
    "insert into select_time (major, stime, etime) values (%d, '%s', '%s');\n"
    % (major_cnt, '2020-5-1 00:00:00', '2020-6-1 00:00:00')
  )  
  major_cnt += 1
f.write('\n')

# Student information
for i in range(student_num):
  grade = rand_grade()
  f.write('''
insert into student (id, name, password, grade, class, major) 
values ('%s', '%s', '%s', '%s', %d, %d);
''' % (rand_id('U', grade, i), rand_name(), rand_pw(), grade, 
  random.randint(1, 15), random.randint(1, len(majors))))  
f.write('\n')

# Faculty information
faculties = []
for i in range(faculty_num):
  faculty = rand_id('T', rand_grade(), i) 
  f.write('''
insert into faculty (id, name, password, position)
values ('%s', '%s', '%s', '%s');
''' % (faculty, rand_name(), rand_pw(), random.choice(['教授', '副教授', '讲师'])))
  faculties.append(faculty)
f.write('\n')

# Course information
for i in range(course_num):
  sdate = random.randint(1, 12)
  edate = sdate + random.randint(4, 8)
  stime, etime = rand_time()
  f.write('''
insert into course (teacher, credit, capacity, sdate, edate, day, stime, etime, valid)
values ('%s', %d, %d, %d, %d, '%s', '%s', '%s', '通过');
''' % (random.choice(faculties), random.randint(1, 8)/2.0, random.randint(3, 20)*10, sdate, edate, 
  random.choice(['周一', '周二', '周三', '周四', '周五', '周六', '周日']), stime, etime))
  
f.write('\n')

f.close()