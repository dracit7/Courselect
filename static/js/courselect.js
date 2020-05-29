
function select_course(id) {
  $.post("/auth/student/select", {
    course: id, 
  });
}

function unselect_course(id) {
  $.post("/auth/student/unselect", {
    course: id, 
  });
}

function list_student(id) {
  window.location.replace("/auth/studentlist?cid=" + id);
}

function delete_course(id) {
  $.post("/auth/faculty/coursedelete", {
    course: id, 
  });
}

function permit_course(id) {
  $.post("/auth/admin/coursepermit", {
    course: id, 
  });
}