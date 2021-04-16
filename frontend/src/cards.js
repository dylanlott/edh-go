/*  cards.js
This file contains functions mapped to Magic the Gathering 
keywords and mechanics. For example, Mill, Draw, Exile, Fetch, 
etc... Fetch is the most powerful and simple of all of them,
and technically any other function could be expressed in terms
of fetch's functionality, but we compose others for ease of use.
*/

// function draw() {
//   const card = this.self.boardstate.Library.shift();
//   this.self.boardstate.Hand.push(card);
//   this.mutateBoardState()
// },

// function mill() {
//   const card = this.self.boardstate.Library.shift();
//   this.self.boardstate.Graveyard.push(card);
//   this.mutateBoardState()
// },

// @param `src` is the source field of cards the target card is in. 
// @param `target` is the card that's being fetched
// @param `dst` is the destination field of the fetched card
// NB: We always want to pass cards around by ID, since we're 
// planning on these being unique.
// @returns: `src`, `dst`
function fetch(src, target, dst) {
    let obj = src.find((v, idx) => {
        if (v.ID === target.ID) {
            console.log(`target found, moving ${target} from ${src} -> ${dst}`)
            src2 = src.splice(1, idx)
            dst2 = dst.push(v)
            console.log(`target moved: src: ${src2} dst: ${dst2}`)
            return src2, dst2
        }
    })
    if (obj === undefined) {
        console.error(`unable to find target ${target}`)
        return src, dst
    }
    // if we get here, we have a weird result. 
    // log and return src.
    console.log('weird, we shouldnt be here', src, dst)
    return src, dst
}
// function increaseLife () {
//   this.self.boardstate.Life++
//   this.mutateBoardState()
// },
// function decreaseLife() {
//   this.self.boardstate.Life--
//   this.mutateBoardState()
// },
// function tap(card) {
//   card.Tapped = !card.Tapped
//   this.mutateBoardState()
// },