
id("forbidden").onclick = function () {
    prompt("TYPE THE PASSWORD TO ACCESS:");
    id("forbtext").style.color = "red";
    alert("ACCESS DENIED!\n\nERR_CODE: SYSTEM_INTRUSER");
}
 