// ua ;
@id("My ID")
permit ( 
    principal == User::"Alice",
    action == Action::"view", // ciao ;
    resource in Album::"adventure"
)
when { principal.age == 18 && 
    context.demoRequest == true }
unless {
  principal has isTerminated && principal.isTerminated
};