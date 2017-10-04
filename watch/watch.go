// package filetree

// import (
// 	"context"
// )

// type watcher interface {
// 	Watch(ctx context.Context, ch chan Change)
// }

// // set by different os-specific drivers
// var newWatcher func(tree *Tree, match *MatcherSet) (watcher, error)

// // Watch the tree for changes in the underlying filesystem
// func (t *Tree) Watch(ctx context.Context, matcher *MatcherSet) (chan Change, error) {
// 	w, err := newWatcher(t, matcher)
// 	if err != nil {
// 		return nil, err
// 	}

// 	changes := make(chan Change)
// 	go w.Watch(ctx, changes)

// 	return changes, nil
// }
