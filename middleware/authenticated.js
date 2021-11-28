export default function({ redirect }) {
  if (
    localStorage.getItem('loggedIn') === 'false' ||
    !localStorage.getItem('loggedIn')
  ) {
    //return redirect('/sign-in')
  }
}
