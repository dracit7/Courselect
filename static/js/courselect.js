
function select_course(id) {
  $.post("/auth/student/select", {
    course: id, 
  });
}