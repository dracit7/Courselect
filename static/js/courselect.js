
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